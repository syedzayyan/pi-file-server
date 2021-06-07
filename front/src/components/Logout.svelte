<script>
    import { auth } from "./auth.js";
    import router from 'page'
    import axios from "axios";
    import { onDestroy } from "svelte";
    let auth_state;
    auth.subscribe((auth) => {
        auth_state = auth;
    });
    const unsubscribe = auth.subscribe((auth) => {
        auth_state = auth;
    });
    onDestroy(unsubscribe);
    import baseAPIURLPath from "./baseAPIPath";
    const logout = () => {
        axios({
            method: "post",
            url: baseAPIURLPath + "user/logout",
            withCredentials: true,
        }).then((res) => {
            console.log(res);
            auth.set(false);
            localStorage.setItem("auth_state", `false ${new Date()}`);
            router.redirect("/")
        }).catch(err => {
            console.log(err)
            localStorage.setItem("auth_state", `false ${new Date()}`);
        });
    };
</script>

<button on:click|preventDefault={logout} class="btn btn-primary">Logout</button>
