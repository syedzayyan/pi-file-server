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
    var name;
    var password;
    var admin_password;
    var errors;
    const signUp = () => {
        var formLogin = new URLSearchParams();
        formLogin.append("email", email);
        formLogin.append("name", name);
        formLogin.append("password", password);
        axios({
            method: "post",
            url: baseAPIURLPath + "user/signup/" + admin_password,
            data: formLogin,
            headers: { "Content-Type": "application/x-www-form-urlencoded" },
            withCredentials: true,
        })
            .then((res) => {
                alert(res.data);
            })
            .catch((err) => {
                email = "";
                name = "";
                password = "";
                admin_password = "";
                errors = err;
            });
    };
</script>

<Modal addText="Sign Up">
    {#if errors}
        <span style="color:red"
            >There is an error. Probably duplicate email {errors}</span
        >
    {/if}
    <form on:submit|preventDefault={signUp}>
        <label for="password">Admin Password:</label>
        <input type="password" name="password" bind:value={admin_password} /><br
        />
        <label for="name">Name:</label>
        <input type="text" name="name" bind:value={name} /><br />
        <label for="email">Email:</label>
        <input type="email" name="email" bind:value={email} /><br />
        <label for="password">Password:</label>
        <input type="password" name="password" bind:value={password} /><br />
        <input type="submit" value="login" class="btn" />
    </form>
</Modal>
