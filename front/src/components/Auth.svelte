<script>
    import { auth } from "./auth.js";
    import { onDestroy } from "svelte";
    import Login from "./Login.svelte";
    import Logout from "./Logout.svelte";
    import SignUp from "./SignUp.svelte";

    let auth_state;
    auth.subscribe((auth) => {
        auth_state = auth;
    });
    const unsubscribe = auth.subscribe((auth) => {
        auth_state = auth;
    });
    onDestroy(unsubscribe);
</script>

<section>
    {#if auth_state}
        <Logout />
    {:else}
        <Login />
        &nbsp;
        <SignUp />
    {/if}
</section>

<style>
    section {
        text-align: center;
        display: flex;
    }
</style>