import { writable } from 'svelte/store';

let auth_state

if (localStorage.getItem("auth_state") === null) {
    localStorage.setItem("auth_state", `false ${new Date()}`)
}


if (localStorage.getItem("auth_state").split(' ')[0] === "true") {
    auth_state = true
}
if (localStorage.getItem("auth_state").split(' ')[0] === "false") {
    auth_state = false
}


let exp = localStorage.getItem("auth_state").split(' ')
exp.shift()
var expTime = new Date(exp.join(' '))
let diff = new Date(new Date - expTime);

if (diff.getHours() > 12) {
    localStorage.setItem("auth_state", `false ${new Date()}`)
    auth_state = false
}

export const auth = writable(auth_state);
