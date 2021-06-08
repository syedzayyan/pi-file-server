<script>
    export let url;
    export let filename;
    export let urlArray;
    export let fileNameArray;
    import Modal from "./Modal.svelte";
    function next() {
        if (urlArray.indexOf(url) === (urlArray.length - 1)) {
            return
        }
        url = urlArray[urlArray.indexOf(url) + 1];
        filename = fileNameArray[fileNameArray.indexOf(filename) + 1]
    }
    function prev() {
        if (urlArray.indexOf(url) === 0) {
            return
        }
        url = urlArray[urlArray.indexOf(url) - 1];
        filename = fileNameArray[fileNameArray.indexOf(filename) - 1]
    }
</script>

<section>
    <Modal addText="Preview ?">
        <div>
            <button class = "btn btn-info" on:click = {prev}>Previous</button>
            {filename}
            <button class = "btn btn-info" on:click = {next}>Next</button>
        </div>
        {#if ["jpg", "jpeg", "gif", "tif", "bmp", "png", "eps", "webp"].includes(filename
                .split(".")
                .pop())}
            <img
                width="100%"
                alt="preview"
                src={url}
            />
        {:else if ["mp3", "wav", "m4a", "flac", "wma", "ogg", "aac"].includes(filename
                .split(".")
                .pop())}
            <!-- svelte-ignore a11y-media-has-caption -->
            <audio width="100%" src={url} controls />
        {:else if ["mp4", "mkv"].includes(filename
                .split(".")
                .pop())}
            <!-- svelte-ignore a11y-media-has-caption -->
            <video width="100%" src={url} controls />
        {:else}
            <span>Can't read file</span>
        {/if}
    </Modal>
</section>
<style>
    div{
        width : 100%;
        display : flex;
        justify-content: space-between;
    }
</style>