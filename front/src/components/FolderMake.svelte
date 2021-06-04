<script>
    import axios from "axios";
    import Modal from "./Modal.svelte";
    var folderName;
    export let baseAPIURLPath 
    export let slug
    function newFolder(e){
        e.preventDefault()
        axios.post(baseAPIURLPath + "folder/" + slug + folderName).then((res) => {
            console.log(res)
            window.location.reload();
        })     
    }
</script>
<Modal addText = {`Add Folder`}>
    <form on:submit = {(e) => {newFolder(e)}} >
        <input class = "add-folder-modal" type = "text" name="folderName" bind:value={folderName}/>
        <input type="submit" value = "+ Folder" class = "btn btn-primary"/>
    </form>
</Modal>

<style>
.add-folder-modal{
    padding:2rem 2rem;
    width: 40%;
}
</style>