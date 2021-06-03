<script>
    import axios from "axios";
    let data;
    let loaded;
    let baseAPIURLPath = "http://localhost:8080/"
    let fileAPIPath = baseAPIURLPath + "fileserver/drive/"
    export let slug;
    $ : {
        axios
            .get(baseAPIURLPath + "files/" + slug)
            .then((res) => {
                data = res.data;
        })
        .then(() => {
            loaded = true;
        });
    }
    function trash(name){
        axios.delete(baseAPIURLPath + "files" + slug + name).then(res => {
            console.log(res);
            window.location.reload();
        })
    }
</script>

<section>
    <table id="files">
        <tr>
            <th>Icon</th>
            <th>File Name</th>
            <th>Delete ?</th>
            <th>File Size</th>
            <th>Last Modified</th>
        </tr>
        {#if loaded}
            {#each Object.entries(data) as [_, value]}
                    {#if value.fileorfolder}
                    <tr>
                        <td><svg xmlns="http://www.w3.org/2000/svg"  viewBox="0 0 64 64" width="64px" height="64px"><linearGradient id="UC11dU6fpD1Whu1sZEZF0a" x1="14.296" x2="14.296" y1="42.833" y2="53.022" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#6dc7ff"/><stop offset="1" stop-color="#e6abff"/></linearGradient><path fill="url(#UC11dU6fpD1Whu1sZEZF0a)" d="M11.702,43.702l6.596,6.596C18.926,50.926,18.481,52,17.593,52H12c-1.105,0-2-0.895-2-2 v-5.593C10,43.519,11.074,43.074,11.702,43.702z"/><linearGradient id="UC11dU6fpD1Whu1sZEZF0b" x1="32" x2="32" y1="8" y2="55.938" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#1a6dff"/><stop offset="1" stop-color="#c822ff"/></linearGradient><path fill="url(#UC11dU6fpD1Whu1sZEZF0b)" d="M55,20h-1v-6c0-1.654-1.346-3-3-3v-1c0-1.103-0.897-2-2-2H15c-1.103,0-2,0.897-2,2v1 c-1.654,0-3,1.346-3,3v6H9c-1.654,0-3,1.346-3,3v6v1v21c0,2.757,2.243,5,5,5h42c2.757,0,5-2.243,5-5V31v-2v-6 C58,21.346,56.654,20,55,20z M55,22c0.552,0,1,0.448,1,1v2.026c-0.584-0.442-1.257-0.773-2-0.925V22H55z M15,10h34v1H15V10z M13,13 h38c0.552,0,1,0.448,1,1v10h-9.929c-1.007,0-1.94,0.5-2.497,1.337l-2.813,4.218C36.575,29.833,36.265,30,35.929,30h-7.857 c-0.336,0-0.646-0.167-0.833-0.446l-2.813-4.216C23.869,24.5,22.936,24,21.929,24H12V14C12,13.448,12.448,13,13,13z M9,22h1v2.101 c-0.743,0.152-1.416,0.482-2,0.925V23C8,22.448,8.448,22,9,22z M56,51c0,1.654-1.346,3-3,3H11c-1.654,0-3-1.346-3-3V30v-1 c0-1.654,1.346-3,3-3h10.929c0.336,0,0.646,0.167,0.833,0.446l2.813,4.216C26.131,31.5,27.064,32,28.071,32h7.857 c1.007,0,1.94-0.5,2.497-1.337l2.813-4.218C41.425,26.167,41.735,26,42.071,26H53c1.654,0,3,1.346,3,3v2V51z"/><linearGradient id="UC11dU6fpD1Whu1sZEZF0c" x1="33" x2="33" y1="8" y2="55.938" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#1a6dff"/><stop offset="1" stop-color="#c822ff"/></linearGradient><path fill="url(#UC11dU6fpD1Whu1sZEZF0c)" d="M21 17H45V19H21z"/><linearGradient id="UC11dU6fpD1Whu1sZEZF0d" x1="49" x2="49" y1="8" y2="55.938" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#1a6dff"/><stop offset="1" stop-color="#c822ff"/></linearGradient><path fill="url(#UC11dU6fpD1Whu1sZEZF0d)" d="M44 50H54V52H44z"/><linearGradient id="UC11dU6fpD1Whu1sZEZF0e" x1="49" x2="49" y1="8" y2="55.938" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#1a6dff"/><stop offset="1" stop-color="#c822ff"/></linearGradient><path fill="url(#UC11dU6fpD1Whu1sZEZF0e)" d="M44 46H54V48H44z"/><linearGradient id="UC11dU6fpD1Whu1sZEZF0f" x1="32" x2="32" y1="8" y2="55.938" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#1a6dff"/><stop offset="1" stop-color="#c822ff"/></linearGradient><path fill="url(#UC11dU6fpD1Whu1sZEZF0f)" d="M27 21H37V23H27z"/><linearGradient id="UC11dU6fpD1Whu1sZEZF0g" x1="32" x2="32" y1="8" y2="55.938" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#1a6dff"/><stop offset="1" stop-color="#c822ff"/></linearGradient><path fill="url(#UC11dU6fpD1Whu1sZEZF0g)" d="M29 25H35V27H29z"/></svg></td>
                        <td><h1><a href = {"/drives/" + slug + value.filename}>{value.filename}</a></h1></td>
                        <td><button on:click = {() => {trash(value.filename)}}>Delete</button></td>
                        <td><h1>{value.filesize}</h1></td>
                        <td><h1>{value.filemodtime}</h1></td>
                    </tr>
                        {:else} 
                    <tr>
                        <td><svg xmlns="http://www.w3.org/2000/svg"  viewBox="0 0 64 64" width="64px" height="64px"><radialGradient id="JgiLTpLYUVYqJAJq8M0ssa" cx="32" cy="32" r="33.002" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#efdcb1"/><stop offset="0" stop-color="#f2e0bb"/><stop offset=".011" stop-color="#f2e0bc"/><stop offset=".362" stop-color="#f9edd2"/><stop offset=".699" stop-color="#fef4df"/><stop offset="1" stop-color="#fff7e4"/></radialGradient><path fill="url(#JgiLTpLYUVYqJAJq8M0ssa)" d="M55.454,45.668c0.5,0.898,1.504,1.337,2.532,1.332c1.75-0.008,3.153,1.483,3.002,3.262 C60.855,51.838,59.434,53,57.852,53L48,53v1H26h-8H7c-1.71,0-3.086-1.431-2.996-3.161C4.089,49.216,5.545,48,7.17,48L8,48 c1.215,0,2.176-1.083,1.973-2.336C9.813,44.681,8.889,44,7.893,44L3,44c-1.71,0-3.086-1.431-2.996-3.161 C0.089,39.216,1.545,38,3.17,38l6.33,0c1.381,0,2.5-1.119,2.5-2.5v0c0-1.381-1.119-2.5-2.5-2.5h0C8.119,33,7,31.881,7,30.5v0 C7,29.119,8.119,28,9.5,28H18v-7h-7.5c-1.995,0-3.601-1.67-3.495-3.688C7.104,15.419,8.803,14,10.698,14L18,14v-4h24h6h8.5 c1.995,0,3.601,1.67,3.495,3.688C59.896,15.581,58.197,17,56.302,17L53.5,17c-1.995,0-3.601,1.67-3.495,3.688 C50.104,22.581,51.803,24,53.698,24L56,24c1.777,0,3.194,1.546,2.978,3.366c-0.179,1.509-1.546,2.572-3.064,2.635 c-1.197,0.05-2.122,1.153-1.874,2.406C54.228,33.355,55.123,34,56.089,34H60.5c1.995,0,3.601,1.67,3.495,3.688 C63.896,39.581,62.197,41,60.302,41L58,41C55.808,41,54.166,43.35,55.454,45.668z M2.5,33L2.5,33C3.881,33,5,31.881,5,30.5v0 C5,29.119,3.881,28,2.5,28h0C1.119,28,0,29.119,0,30.5v0C0,31.881,1.119,33,2.5,33z M2.5,33L2.5,33C3.881,33,5,31.881,5,30.5v0 C5,29.119,3.881,28,2.5,28h0C1.119,28,0,29.119,0,30.5v0C0,31.881,1.119,33,2.5,33z"/><linearGradient id="JgiLTpLYUVYqJAJq8M0ssb" x1="33" x2="33" y1="58" y2="6" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#41bfec"/><stop offset=".232" stop-color="#4cc5ef"/><stop offset=".644" stop-color="#6bd4f6"/><stop offset="1" stop-color="#8ae4fd"/></linearGradient><path fill="url(#JgiLTpLYUVYqJAJq8M0ssb)" d="M50,58H16c-1.657,0-3-1.343-3-3V9c0-1.657,1.343-3,3-3h20.757 c0.796,0,1.559,0.316,2.121,0.879l13.243,13.243C52.684,20.684,53,21.447,53,22.243V55C53,56.657,51.657,58,50,58z"/><linearGradient id="JgiLTpLYUVYqJAJq8M0ssc" x1="38.879" x2="46.454" y1="20.121" y2="12.546" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#c6effd"/><stop offset=".375" stop-color="#b7ecfd"/><stop offset="1" stop-color="#95e6fd"/></linearGradient><path fill="url(#JgiLTpLYUVYqJAJq8M0ssc)" d="M52.121,20.122L38.878,6.879c-0.255-0.255-0.556-0.452-0.878-0.6V18c0,1.657,1.343,3,3,3 h11.721C52.574,20.678,52.377,20.377,52.121,20.122z"/><linearGradient id="JgiLTpLYUVYqJAJq8M0ssd" x1="45.5" x2="45.5" y1="24.083" y2="18.083" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#42c6ee"/><stop offset=".48" stop-color="#3bc3ed"/><stop offset="1" stop-color="#2ebeea"/></linearGradient><path fill="url(#JgiLTpLYUVYqJAJq8M0ssd)" d="M41,21c-1.657,0-3-1.343-3-3v3c0,1.657,1.343,3,3,3h12v-1.757 c0-0.434-0.102-0.855-0.279-1.243H41z"/><linearGradient id="JgiLTpLYUVYqJAJq8M0sse" x1="22" x2="22" y1="21" y2="6" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#81dcf7"/><stop offset=".48" stop-color="#8ce1f9"/><stop offset="1" stop-color="#9ee8fd"/></linearGradient><path fill="url(#JgiLTpLYUVYqJAJq8M0sse)" d="M21.5,21H13V9c0-1.657,1.343-3,3-3h9.5C26.881,6,28,7.119,28,8.5v0 c0,1.381-1.119,2.5-2.5,2.5h-5c-1.381,0-2.5,1.119-2.5,2.5v0c0,1.381,1.119,2.5,2.5,2.5h1c1.381,0,2.5,1.119,2.5,2.5v0 C24,19.881,22.881,21,21.5,21z M29,17c-1.105,0-2,0.895-2,2s0.895,2,2,2c1.105,0,2-0.895,2-2S30.105,17,29,17z"/><linearGradient id="JgiLTpLYUVYqJAJq8M0ssf" x1="41.5" x2="41.5" y1="41" y2="58" gradientUnits="userSpaceOnUse" spreadMethod="reflect"><stop offset="0" stop-color="#47c3ee"/><stop offset=".98" stop-color="#32b5e8"/><stop offset="1" stop-color="#32b5e8"/></linearGradient><path fill="url(#JgiLTpLYUVYqJAJq8M0ssf)" d="M44,41l9,0v14c0,1.657-1.343,3-3,3H33c-1.657,0-3-1.343-3-3v0c0-1.657,1.343-3,3-3l11.5,0 c1.381,0,2.5-1.119,2.5-2.5v0c0-1.381-1.119-2.5-2.5-2.5H44c-1.657,0-3-1.343-3-3v0C41,42.343,42.343,41,44,41z M36,41 c-1.657,0-3,1.343-3,3s1.343,3,3,3s3-1.343,3-3S37.657,41,36,41z"/></svg></td>
                        <td><a href = {fileAPIPath + slug + value.filename}><h1>{value.filename}</h1></a></td>
                        <td><button on:click = {() => {trash(value.filename)}}>Delete</button></td>
                        <td><h1>{value.filesize}</h1></td>
                        <td><h1>{value.filemodtime}</h1></td>
                    </tr>
                    {/if}
            {/each}
        {/if}
    </table>
</section>

<style>
    table {
    background: white;
    border: 1px solid #E4E5E7;
    border-collapse: collapse;
    width: 100%;
}

td, th {
    text-align: left;
    padding: 9px 18px;
}

th:last-child, td:last-child {
    text-align: right;
    color: #6A6C6F;
}

tr {
    border-bottom: 1px solid #E4E5E7;
}

tr:nth-child(2n) {
    background-color: #F7F9FA;
}



</style>
