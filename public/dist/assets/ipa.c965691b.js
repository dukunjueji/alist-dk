import{d as c,e as r,j as e,a4 as p,B as n,bk as u,cB as g,o as a,cA as f}from"./index.d66f6448.js";import{a as d}from"./useUtil.9a317e9f.js";import{F as h}from"./File.cbcdeeba.js";import"./api.7a22b4b8.js";import"./icon.e912e56e.js";import"./index.1733bd32.js";import"./index.0d244616.js";import"./Layout.469b2801.js";import"./EncodingSelect.1b4732a4.js";import"./index.c106e327.js";import"./FolderTree.5cc6857b.js";const U=()=>{const t=c(),[o,i]=r(!1),[s,l]=r(!1),{currentObjLink:m}=d();return e(h,{get children(){return e(p,{spacing:"$2",get children(){return[e(n,{as:"a",get href(){return`itms-services://?action=download-manifest&url=${u}/i/${g(encodeURIComponent(a.raw_url)+"/"+f(encodeURIComponent(a.obj.name)))}.plist`},onClick:()=>{i(!0)},get children(){return t(`home.preview.${o()?"installing":"install"}`)}}),e(n,{as:"a",colorScheme:"primary",get href(){return"apple-magnifier://install?url="+encodeURIComponent(m(!0))},onClick:()=>{l(!0)},get children(){return t(`home.preview.${s()?"tr-installing":"tr-install"}`)}})]}})}})};export{U as default};