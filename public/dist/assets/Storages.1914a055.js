import{d as b,j as e,N as R,aa as X,a4 as o,aQ as v,bM as j,al as A,Z as T,X as x,ck as l,ch as p,u as D,a as _,bp as c,B as i,aO as F,br as h,n as O,e as k,w as G,bI as N,r as P,b4 as W,b5 as Q,b6 as U,b7 as Z,b8 as q,b9 as z,ba as E,x as u,bb as J,bc as K,bd as Y,c5 as ee,U as re,V as y,c6 as te,cf as ae,cg as ne,ci as C,cj as se}from"./index.590d4491.js";import{o as le}from"./index.a229a717.js";import{D as oe}from"./DeletePopover.3542c434.js";function B(r){const a=b(),{to:g}=D(),[m,d]=_(()=>c.post(`/admin/storage/delete?id=${r.storage.id}`)),[S,n]=_(()=>c.post(`/admin/storage/${r.storage.disabled?"enable":"disable"}?id=${r.storage.id}`));return[e(i,{onClick:()=>{g(`/@manage/storages/edit/${r.storage.id}`)},get children(){return a("global.edit")}}),e(i,{get loading(){return S()},get colorScheme(){return r.storage.disabled?"success":"warning"},onClick:async()=>{const s=await n();F(s,()=>{r.refresh()})},get children(){return a(`global.${r.storage.disabled?"enable":"disable"}`)}}),e(oe,{get name(){return r.storage.mount_path},get loading(){return m()},onClick:async()=>{const s=await d();h(s,()=>{O.success(a("global.delete_success")),r.refresh()})}})]}function ce(r){const a=b();return e(x,{w:"$full",spacing:"$2",rounded:"$lg",border:"1px solid $neutral7",get background(){return R("$neutral2","$neutral3")()},p:"$3",get _hover(){return{border:`1px solid ${X()}`}},get children(){return[e(o,{spacing:"$2",get children(){return[e(v,{fontWeight:"$medium",css:{wordBreak:"break-all"},get children(){return r.storage.mount_path}}),e(j,{colorScheme:"info",get children(){return a(`drivers.drivers.${r.storage.driver}`)}})]}}),e(o,{get children(){return[e(v,{get children(){return[A(()=>a("storages.common.status")),":\xA0"]}}),e(T,{css:{wordBreak:"break-all"},overflowX:"auto",get innerHTML(){return r.storage.status}})]}}),e(v,{css:{wordBreak:"break-all"},get children(){return r.storage.remark}}),e(o,{spacing:"$2",get children(){return e(B,r)}})]}})}function ie(r){const a=b();return e(p,{get children(){return[e(l,{get children(){return r.storage.mount_path}}),e(l,{get children(){return a(`drivers.drivers.${r.storage.driver}`)}}),e(l,{get children(){return r.storage.order}}),e(l,{get children(){return r.storage.status}}),e(l,{get children(){return r.storage.remark}}),e(l,{get children(){return e(o,{spacing:"$2",get children(){return e(B,r)}})}})]}})}const be=()=>{const r=b();le("manage.sidemenu.storages");const{to:a}=D(),[g,m]=_(()=>c.get("/admin/storage/list")),[d,S]=k([]),n=async()=>{const t=await m();h(t,w=>S(w.content))},[s,L]=k([]),[f,M]=k([]);(async()=>{const t=await c.get("/admin/driver/names");h(t,w=>L(w))})(),n();const I=async()=>{const t=await c.post("/admin/storage/load_all");h(t,()=>{O.success(r("storages.other.start_load_success"))})},V=G(()=>d().filter(t=>f().length===0?!0:f().includes(t.driver))),[$,H]=N("storages-layout","grid");return e(x,{spacing:"$3",alignItems:"start",w:"$full",get children(){return[e(o,{spacing:"$2",gap:"$2",w:"$full",wrap:{"@initial":"wrap","@md":"unset"},get children(){return[e(i,{colorScheme:"accent",get loading(){return g()},onClick:n,get children(){return r("global.refresh")}}),e(i,{onClick:()=>{a("/@manage/storages/add")},get children(){return r("global.add")}}),e(i,{colorScheme:"warning",get loading(){return g()},onClick:I,get children(){return r("storages.other.load_all")}}),e(P,{get when(){return s().length>0},get children(){return e(W,{multiple:!0,get value(){return f()},onChange:M,get children(){return[e(Q,{get children(){return[e(U,{get children(){return r("storages.other.filter_by_driver")}}),e(Z,{}),e(q,{})]}}),e(z,{get children(){return e(E,{get children(){return e(u,{get each(){return s()},children:t=>e(J,{value:t,get children(){return[e(K,{get children(){return r(`drivers.drivers.${t}`)}}),e(Y,{})]}})})}})}})]}})}}),e(ee,{get checked(){return $()==="table"},onChange:t=>{H(t.currentTarget.checked?"table":"grid")},get children(){return r("storages.other.table_layout")}})]}}),e(re,{get children(){return[e(y,{get when(){return $()==="grid"},get children(){return e(te,{w:"$full",gap:"$2_5",templateColumns:{"@initial":"1fr","@lg":"repeat(auto-fill, minmax(324px, 1fr))"},get children(){return e(u,{get each(){return V()},children:t=>e(ce,{storage:t,refresh:n})})}})}}),e(y,{get when(){return $()==="table"},get children(){return e(T,{w:"$full",overflowX:"auto",get children(){return e(ae,{highlightOnHover:!0,dense:!0,get children(){return[e(ne,{get children(){return e(p,{get children(){return[e(u,{each:["mount_path","driver","order","status","remark"],children:t=>e(C,{get children(){return r(`storages.common.${t}`)}})}),e(C,{get children(){return r("global.operations")}})]}})}}),e(se,{get children(){return e(u,{get each(){return d()},children:t=>e(ie,{storage:t,refresh:n})})}})]}})}})}})]}})]}})};export{be as default};