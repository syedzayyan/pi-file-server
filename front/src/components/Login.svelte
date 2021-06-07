<script>
    import Modal from "./Modal.svelte";
    import { auth } from "./auth.js";
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
    var email;
    var password;
    var errors;
    const login = () => {
        var formLogin = new URLSearchParams();
        formLogin.append("email", email);
        formLogin.append("password", password);
        axios({
            method: "post",
            url: baseAPIURLPath + "user/login",
            data: formLogin,
            headers: { "Content-Type": "application/x-www-form-urlencoded" },
            withCredentials: true,
        })
            .then((res) => {
                console.log(res);
                auth.set(true);
                localStorage.setItem("auth_state", `true ${new Date()}`);
            })
            .catch((err) => {
                email = "";
                password = "";
                errors = err;
            });
    };
</script>

<Modal addText="Login">
    {#if errors}
        <span style="color:red">Invalid Creds {errors}</span>
    {/if}
    <form on:submit|preventDefault={login}>
        <label for="email">Email:</label>
        <input type="email" name="email" bind:value={email} /><br />
        <label for="password">Password:</label>
        <input type="password" name="password" bind:value={password} /><br />
        <input type="submit" value="login" class="btn" />
    </form>
</Modal>
