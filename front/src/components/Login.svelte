<script>
    import Modal from "./Modal.svelte";
    import { auth } from "./auth.js";
    import axios from "axios";
    import { onDestroy } from "svelte";
    let auth_state;
    auth.subscribe(auth => {
		auth_state = auth;
	});
    const unsubscribe = auth.subscribe(auth => {
	    auth_state = auth;
    });
    onDestroy(unsubscribe)
    import baseAPIURLPath from "./baseAPIPath"
    var email;
    var password;
    var formLogin = new URLSearchParams();
    formLogin.append("email", "masudzayyan@gmail.com");
    formLogin.append("password", "Bijoy50;");
    const login = () => {
        axios({
            method: "post",
            url: baseAPIURLPath + "user/login",
            data: formLogin,
            headers: {'Content-Type': 'application/x-www-form-urlencoded'},
            withCredentials : true,
        }).then(res => {
            console.log(res)
            auth.set(true)
            localStorage.setItem("auth_state", `true ${new Date()}`)
        });
    };
    const logout = () => {
        axios({
            method: "post",
            url: baseAPIURLPath + "user/logout",
            withCredentials : true,
        }).then(res => {
            console.log(res)
            auth.set(false)
            localStorage.setItem("auth_state", `false ${new Date()}`)
        });
    }
</script>

<section>
    {#if auth_state}
        <button on:click|preventDefault = {logout} class = "btn btn-primary">Logout</button>
    {:else}
    <Modal addText="Login">
        <form on:submit|preventDefault={login}>
            <label for="email">Email:</label>
            <input type="email" name="email" bind:value={email} />
            <label for="password">Password:</label>
            <input type="password" name="password" bind:value={password} />
            <input type="submit" value="login" class="btn" />
        </form>
    </Modal>
    {/if}
</section>

<style>
    section {
        text-align: center;
    }
</style>
