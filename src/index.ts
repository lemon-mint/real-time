import KalmanFilter from 'kalmanjs';

const timeH1 = document.getElementById("time");
const errorPre = document.getElementById("error");
const errordiv = document.getElementById("error-div");
const stat = document.getElementById("status");


let timeOffset: number = 0;
let errorString: string = "";
let DOMTimeOff: number = new Date().getTime() - performance.now();
let displayTime: string = new Date().toLocaleString();

let statusText: string = "Please wait, synchronizing...";

function getCurrentTime() {
  return performance.now() + DOMTimeOff + timeOffset;
}

function updateTime() {
  const time = getCurrentTime();
  const date = new Date(time);
  displayTime = date.toLocaleString();
  if (timeH1.textContent !== displayTime) {
    timeH1.textContent = displayTime;
  }

  if (errorString) {
    errordiv.hidden = false;
    errorPre.innerText = errorString;
  } else if (!errordiv.hidden) {
    errordiv.hidden = true;
  }

  if (stat.textContent !== statusText) {
    stat.textContent = statusText;
  }

  requestAnimationFrame(updateTime);
}

requestAnimationFrame(updateTime);

function setTimeOffset(offset: number) {
  timeOffset = timeOffset + offset;
}

async function getServerOff(): Promise<number> {
  const t0 = getCurrentTime();
  const response = await fetch("/time");
  const t3 = getCurrentTime();
  const json = await response.json();
  const t1 = json.t1;
  const t2 = json.t2;

  const offset = (t1 - t0 + t2 - t3) / 2;
  return offset;
}

async function sleep(t: number) {
  return new Promise((resolve) => setTimeout(resolve, t));
}

async function syncTime(repeat: boolean = false, multiplier: number = 1.0, samples: number = 10) {
  await getServerOff(); // Preheat

  let offset = 0;
  const TIME_SYNC_MEASURE_COUNT = samples;
  const kf = new KalmanFilter();
  for (let i = 0; i < TIME_SYNC_MEASURE_COUNT; i++) {
    console.log("Measurement: " + i);
    console.log("Requesting time offset...");
    const off = await getServerOff();
    const kfOff = kf.filter(off);
    console.log("Server offset: " + off);
    //console.log("Kalman filter offset: " + kfOff);
    offset += kfOff;
    statusText = "Synchronizing... " + (offset / (i + 1)).toFixed(5) + "ms";
    await sleep(20);
  }

  offset /= TIME_SYNC_MEASURE_COUNT;
  offset *= multiplier;

  console.log("Synchronized time offset: " + offset);
  setTimeOffset(offset);
  statusText = "" + (getCurrentTime() - new Date().getTime()).toFixed(5) + "ms";
  if (repeat) {
    setTimeout(async () => {
      await syncTime(true, multiplier);
    }, 1000 * 30);
    console.log(
      "Next Synchronization is scheduled on " +
        new Date(getCurrentTime() + 1000 * 30).toLocaleString()
    );
  }
}

async function main() {
  try {
    errorString = "";
    statusText = "Please wait, synchronizing...";
    console.log("Starting synchronization...");
    await getServerOff();
    await sleep(500);
    await syncTime(true, 1.0, 20);
    console.log("Successfully synchronized time.");
  } catch (e) {
    console.error(e);
    errorString = e.toString();
    statusText = "Error occurred. Restarting in 5 seconds.";
    setTimeout(main, 5000);
  }
}

main().then();
