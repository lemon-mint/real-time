function t(t){return t&&t.__esModule?t.default:t}var e;function n(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}function i(t,e){for(var n=0;n<e.length;n++){var i=e[n];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(t,i.key,i)}}e=function(){function t(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},i=e.R,o=void 0===i?1:i,s=e.Q,r=void 0===s?1:s,a=e.A,c=void 0===a?1:a,u=e.B,h=void 0===u?0:u,l=e.C,f=void 0===l?1:l;n(this,t),this.R=o,this.Q=r,this.A=c,this.C=f,this.B=h,this.cov=NaN,this.x=NaN}var e,o,s;
/**
* KalmanFilter
* @class
* @author Wouter Bulten
* @see {@link http://github.com/wouterbulten/kalmanjs}
* @version Version: 1.0.0-beta
* @copyright Copyright 2015-2018 Wouter Bulten
* @license MIT License
* @preserve
*/return e=t,o=[{key:"filter",value:function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:0;if(isNaN(this.x))this.x=1/this.C*t,this.cov=1/this.C*this.Q*(1/this.C);else{var n=this.predict(e),i=this.uncertainty(),o=i*this.C*(1/(this.C*i*this.C+this.Q));this.x=n+o*(t-this.C*n),this.cov=i-o*this.C*i}return this.x}},{key:"predict",value:function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:0;return this.A*this.x+this.B*t}},{key:"uncertainty",value:function(){return this.A*this.cov*this.A+this.R}},{key:"lastMeasurement",value:function(){return this.x}},{key:"setMeasurementNoise",value:function(t){this.Q=t}},{key:"setProcessNoise",value:function(t){this.R=t}}],o&&i(e.prototype,o),s&&i(e,s),t}();const o=document.getElementById("time"),s=document.getElementById("error"),r=document.getElementById("error-div"),a=document.getElementById("status");let c=0,u="",h=(new Date).getTime()-performance.now(),l=(new Date).toLocaleString(),f="Please wait, synchronizing...";function d(){return performance.now()+h+c}async function m(){const t=d(),e=await fetch("/time"),n=d(),i=await e.json();return(i.t1-t+i.t2-n)/2}async function y(t){return new Promise((e=>setTimeout(e,t)))}async function v(n=!1,i=1,o=10){await m();let s=0;const r=o,a=new(t(e));for(let t=0;t<r;t++){console.log("Measurement: "+t),console.log("Requesting time offset...");const e=await m(),n=a.filter(e);console.log("Server offset: "+e),s+=n,f="Synchronizing... "+(s/(t+1)).toFixed(5)+"ms",await y(20)}s/=r,s*=i,console.log("Synchronized time offset: "+s),function(t){c+=t}(s),f=(d()-(new Date).getTime()).toFixed(5)+"ms",n&&(setTimeout((async()=>{await v(!0,i)}),3e4),console.log("Next Synchronization is scheduled on "+new Date(d()+3e4).toLocaleString()))}requestAnimationFrame((function t(){const e=d(),n=new Date(e);l=n.toLocaleString(),o.textContent!==l&&(o.textContent=l),u?(r.hidden=!1,s.innerText=u):r.hidden||(r.hidden=!0),a.textContent!==f&&(a.textContent=f),requestAnimationFrame(t)})),async function t(){try{u="",f="Please wait, synchronizing...",console.log("Starting synchronization..."),await m(),await y(500),await v(!0,1,20),console.log("Successfully synchronized time.")}catch(e){console.error(e),u=e.toString(),f="Error occurred. Restarting in 5 seconds.",setTimeout(t,5e3)}}().then();
//# sourceMappingURL=index.3fb7c888.js.map
