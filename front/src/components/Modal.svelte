<script>
    export let addText;
    var modal;
    function closeModal() {
        modal.style.display = "none";
    }
    function openModal() {
        modal.style.display = "block";
    }
</script>

<div>
    <button on:click={() => openModal()} class="btn btn-primary">
        {addText}
    </button>
    <!-- The Modal -->
    <div id="myModal" bind:this={modal} class="modal">
        <div class="modal-content">
            <span class="close" on:click={() => closeModal()}>&times;</span>
            <h2>{addText}</h2>
            <slot>
                <h1>Nothing Provided</h1>
            </slot>
        </div>
    </div>
</div>

<style>
    .modal {
        display: none; /* Hidden by default */
        position: fixed; /* Stay in place */
        z-index: 1; /* Sit on top */
        left: 0;
        top: 0;
        width: 100%; /* Full width */
        height: 100%; /* Full height */
        overflow: auto; /* Enable scroll if needed */
        background-color: rgb(0, 0, 0); /* Fallback color */
        background-color: rgba(0, 0, 0, 0.4); /* Black w/ opacity */
    }
    /* Modal Content/Box */
    .modal-content {
        background-color: #fefefe;
        margin: 15% auto; /* 15% from the top and centered */
        padding: 2%;
        border: 1px solid #888;
        width: 60%;
    }
    /* The Close Button */
    .close {
        color: #aaa;
        float: right;
        font-size: 28px;
        font-weight: bold;
    }
    .close:hover,
    .close:focus {
        color: black;
        text-decoration: none;
        cursor: pointer;
    }
</style>
