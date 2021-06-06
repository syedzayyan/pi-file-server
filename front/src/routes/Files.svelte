<script>
    import FileBrowser from "../components/FileBrowser.svelte";
    import FolderMake from "../components/FolderMake.svelte";
    import FileMake from "../components/FileMake.svelte";
    export let params;
    let slug;
    import baseAPIURLPath from "../components/baseAPIPath"
    let fileAPIPath = baseAPIURLPath + "fileserver/drive/"
    $: {
        if (params === undefined) {
            slug = "";
        } else {
            slug = params["0"] + "/";
        }
    }
    import { onDestroy } from "svelte";
    import {auth} from "../components/auth"
    let auth_state;
    auth.subscribe((auth) => {
        auth_state = auth;
    });
    const unsubscribe = auth.subscribe((auth) => {
        auth_state = auth;
    });
    onDestroy(unsubscribe);
</script>

<main>
    {#if auth_state}
    <div>
        &nbsp;
        <FolderMake {slug} {baseAPIURLPath}/>&nbsp;
        <FileMake {slug} {baseAPIURLPath}/>
    </div>
    {:else}
    <span>Login</span>
    {/if}
    <FileBrowser {slug} {baseAPIURLPath} {fileAPIPath}/>
</main>

<style>
    div{
        height : 10vh;
        display : flex;
        align-items: center;
    }
</style>
