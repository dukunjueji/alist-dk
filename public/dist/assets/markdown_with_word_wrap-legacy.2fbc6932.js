!function(){function t(t,n){return function(t){if(Array.isArray(t))return t}(t)||function(t,r){var n=null==t?null:"undefined"!=typeof Symbol&&t[Symbol.iterator]||t["@@iterator"];if(null==n)return;var e,o,u=[],a=!0,i=!1;try{for(n=n.call(t);!(a=(e=n.next()).done)&&(u.push(e.value),!r||u.length!==r);a=!0);}catch(l){i=!0,o=l}finally{try{a||null==n.return||n.return()}finally{if(i)throw o}}return u}(t,n)||function(t,n){if(!t)return;if("string"==typeof t)return r(t,n);var e=Object.prototype.toString.call(t).slice(8,-1);"Object"===e&&t.constructor&&(e=t.constructor.name);if("Map"===e||"Set"===e)return Array.from(t);if("Arguments"===e||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(e))return r(t,n)}(t,n)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function r(t,r){(null==r||r>t.length)&&(r=t.length);for(var n=0,e=new Array(r);n<r;n++)e[n]=t[n];return e}System.register(["./index-legacy.c8b63b33.js","./useUtil-legacy.0da7f9ed.js","./EncodingSelect-legacy.b47a84ac.js","./api-legacy.bfb34aa3.js"],(function(r){"use strict";var n,e,o,u;return{setters:[function(t){n=t.j,e=t.$},function(t){o=t.e},function(t){u=t.M},function(){}],execute:function(){r("default",(function(){var r=t(o(),1)[0];return n(e,{get loading(){return r.loading},get children(){return n(u,{class:"word-wrap",get children(){var t;return null===(t=r())||void 0===t?void 0:t.content}})}})}))}}}))}();
