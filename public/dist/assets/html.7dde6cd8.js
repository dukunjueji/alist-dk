import{j as e,$ as i,e as s,bW as a,r as l,c0 as d}from"./index.590d4491.js";import{e as u,c as g}from"./useUtil.129c3c9a.js";import{E as h}from"./EncodingSelect.b0ad3407.js";import"./api.1e99203c.js";function m(n){const[t,r]=s("utf-8"),{isString:o,text:c}=g(n.children);return e(d,{w:"$full",h:"70vh",pos:"relative",get children(){return[e(a.iframe,{w:"$full",h:"$full",rounded:"$lg",shadow:"$md",get srcdoc(){return c(t())}}),e(l,{when:!o,get children(){return e(h,{get encoding(){return t()},setEncoding:r})}})]}})}const x=()=>{const[n]=u();return e(i,{get loading(){return n.loading},get children(){return e(m,{get children(){var t;return(t=n())==null?void 0:t.content}})}})};export{x as default};