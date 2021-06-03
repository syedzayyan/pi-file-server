<script>
    export let slug;
    export let baseAPIURLPath;
    import axios from "axios";
    import SomeModal from "./Modal.svelte";
    var visible = "hidden";
    var percentCompleted;
    function sendFile() {
        var file = document.querySelector("#file");
        visible = "visible";
        const sendFile = new FormData();
        for (var i = 0; i < file.files.length; i++) {
            sendFile.append("files", file.files[i]);
        }
        axios
            .post(baseAPIURLPath + "files/" + slug, sendFile, {
                headers: {
                    "Content-Type": "multipart/form-data",
                },
                onUploadProgress: (e) => {
                    percentCompleted = Math.round((e.loaded * 100) / e.total);
                },
            })
            .then((res) => {
                console.log(res.data);
                window.location.reload();
            });
    }
</script>

<SomeModal addText={`Add Files`}>
    <div>
        <form on:submit|preventDefault={sendFile}>
            <input
                type="file"
                id="file"
                name="myFile"
                multiple
                class="dark-button"
            />
            <input type="submit" value="upload" class="dark-button" />
            <div style={`visibility:${visible}`}>
                <progress
                    id="file"
                    value={percentCompleted ? percentCompleted : 100}
                    max="100"
                />{percentCompleted ? percentCompleted : 100} % uploaded
            </div>
        </form>
    </div>
</SomeModal>

<style>
    progress {
        display: block; /* default: inline-block */
        width: 80%;
        margin: 2em auto;
        padding: 4px;
        border: 0 none;
        background: #444;
        border-radius: 14px;
        box-shadow: inset 0px 1px 1px rgba(0, 0, 0, 0.5),
            0px 1px 0px rgba(255, 255, 255, 0.2);
    }
    progress::-moz-progress-bar {
        border-radius: 12px;
        background: #fff;
        box-shadow: inset 0 -2px 4px rgba(0, 0, 0, 0.4),
            0 2px 5px 0px rgba(0, 0, 0, 0.3);
    }
    /* webkit */
    @media screen and (-webkit-min-device-pixel-ratio: 0) {
        progress {
            height: 25px;
        }
    }
    progress::-webkit-progress-bar {
        background: transparent;
    }
    progress::-webkit-progress-value {
        border-radius: 12px;
        background: #fff;
        box-shadow: inset 0 -2px 4px rgba(0, 0, 0, 0.4),
            0 2px 5px 0px rgba(0, 0, 0, 0.3);
    }
    /* environnement styles */
</style>