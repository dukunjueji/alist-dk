!function(){function n(n,t){return function(n){if(Array.isArray(n))return n}(n)||function(n,e){var t=null==n?null:"undefined"!=typeof Symbol&&n[Symbol.iterator]||n["@@iterator"];if(null==t)return;var r,a,c=[],i=!0,o=!1;try{for(t=t.call(n);!(i=(r=t.next()).done)&&(c.push(r.value),!e||c.length!==e);i=!0);}catch(l){o=!0,a=l}finally{try{i||null==t.return||t.return()}finally{if(o)throw a}}return c}(n,t)||function(n,t){if(!n)return;if("string"==typeof n)return e(n,t);var r=Object.prototype.toString.call(n).slice(8,-1);"Object"===r&&n.constructor&&(r=n.constructor.name);if("Map"===r||"Set"===r)return Array.from(n);if("Arguments"===r||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))return e(n,t)}(n,t)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function e(n,e){(null==e||e>n.length)&&(e=n.length);for(var t=0,r=new Array(e);t<e;t++)r[t]=n[t];return r}System.register(["./index-legacy.c8b63b33.js","./useUtil-legacy.0da7f9ed.js","./File-legacy.a0828d23.js","./api-legacy.bfb34aa3.js","./icon-legacy.33e39b8a.js","./index-legacy.b95b9a1c.js","./index-legacy.f950bf6a.js","./Layout-legacy.2f8c9cb7.js","./EncodingSelect-legacy.b47a84ac.js","./index-legacy.8620a2aa.js","./FolderTree-legacy.a6a387aa.js"],(function(e){"use strict";var t,r,a,c,i,o,l,u,f,s,y;return{setters:[function(n){t=n.d,r=n.e,a=n.j,c=n.a4,i=n.B,o=n.bm,l=n.cA,u=n.o,f=n.cz},function(n){s=n.a},function(n){y=n.F},function(){},function(){},function(){},function(){},function(){},function(){},function(){},function(){}],execute:function(){e("default",(function(){var e=t(),g=n(r(!1),2),d=g[0],m=g[1],b=n(r(!1),2),p=b[0],h=b[1],j=s().currentObjLink;return a(y,{get children(){return a(c,{spacing:"$2",get children(){return[a(i,{as:"a",get href(){return"itms-services://?action=download-manifest&url="+"".concat(o,"/i/").concat(l(encodeURIComponent(u.raw_url)+"/"+f(encodeURIComponent(u.obj.name))),".plist")},onClick:function(){m(!0)},get children(){return e("home.preview.".concat(d()?"installing":"install"))}}),a(i,{as:"a",colorScheme:"primary",get href(){return"apple-magnifier://install?url="+encodeURIComponent(j(!0))},onClick:function(){h(!0)},get children(){return e("home.preview.".concat(p()?"tr-installing":"tr-install"))}})]}})}})}))}}}))}();
