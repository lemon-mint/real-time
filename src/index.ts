const timeH1 = document.getElementById("time");
const errorPre = document.getElementById("error");
const errordiv = document.getElementById("error-div");
const stat = document.getElementById("status");

let timeOffset: number = 0;
let errorString: string = "";
let DOMTimeOff: number = new Date().getTime() - performance.now();
let displayTime: string = new Date().toLocaleString();

let statusText: string = "잠시 기다려주세요, 동기화 중입니다.";

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

async function syncTime(repeat: boolean = false) {
  let offset = await getServerOff();
  setTimeOffset(offset);
  statusText = "서버와의 시간 차이: " + offset.toFixed(3) + "ms";
  await sleep(1000);
  offset = await getServerOff();
  setTimeOffset(offset);
  statusText = "서버와의 시간 차이: " + offset.toFixed(3) + "ms";
  if (repeat) {
    setTimeout(syncTime, 1000 * 30);
  }
}

async function main() {
  await getServerOff();
  await sleep(500);
  await syncTime(false);
  await sleep(500);
  await syncTime(true);
}

main().then();
