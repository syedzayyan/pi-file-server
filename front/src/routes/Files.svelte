<script>
    import FileBrowser from "../components/FileBrowser.svelte";
    import FolderMake from "../components/FolderMake.svelte";
    import FileMake from "../components/FileMake.svelte";
    export let params;
    let slug;
    let baseAPIURLPath = "http://localhost:8080/"
    let fileAPIPath = baseAPIURLPath + "fileserver/drive/"
    $: {
        if (params === undefined) {
            slug = "";
        } else {
            slug = params["0"] + "/";
        }
    }
</script>

<main>
    <div>
        &nbsp;
        <FolderMake {slug} {baseAPIURLPath}/>&nbsp;
        <FileMake {slug} {baseAPIURLPath}/>
    </div>
    <FileBrowser {slug} {baseAPIURLPath} {fileAPIPath}/>
</main>

<style>
    div{
        height : 10vh;
        display : flex;
        align-items: center;
    }
</style>
