<script>
import Login from "../components/Login.svelte";

</script>
<section>
    <header />
    <nav>
        <div>
            <a href="/drives/">Home</a>
            <a href="/">About</a>
        </div>
        <div>
                <Login />
        </div>
    </nav>
</section>
<slot />

<style>
    nav {
        padding: 2px calc((100% - 800px) / 2) 0;
    }
    a {
        color: #62cb31;
        text-decoration: none;
    }

    a:hover {
        color: #4eb722;
        text-decoration: underline;
    }

    header {
        background-image: -webkit-linear-gradient(
            left,
            #34495e,
            #34495e 25%,
            #9b59b6 25%,
            #9b59b6 35%,
            #3498db 35%,
            #3498db 45%,
            #62cb31 45%,
            #62cb31 55%,
            #ffb606 55%,
            #ffb606 65%,
            #e67e22 65%,
            #e67e22 75%,
            #e74c3c 85%,
            #e74c3c 85%,
            #c0392b 85%,
            #c0392b 100%
        );
        background-image: -moz-linear-gradient(
            left,
            #34495e,
            #34495e 25%,
            #9b59b6 25%,
            #9b59b6 35%,
            #3498db 35%,
            #3498db 45%,
            #62cb31 45%,
            #62cb31 55%,
            #ffb606 55%,
            #ffb606 65%,
            #e67e22 65%,
            #e67e22 75%,
            #e74c3c 85%,
            #e74c3c 85%,
            #c0392b 85%,
            #c0392b 100%
        );
        background-image: -ms-linear-gradient(
            left,
            #34495e,
            #34495e 25%,
            #9b59b6 25%,
            #9b59b6 35%,
            #3498db 35%,
            #3498db 45%,
            #62cb31 45%,
            #62cb31 55%,
            #ffb606 55%,
            #ffb606 65%,
            #e67e22 65%,
            #e67e22 75%,
            #e74c3c 85%,
            #e74c3c 85%,
            #c0392b 85%,
            #c0392b 100%
        );
        background-image: linear-gradient(
            to right,
            #34495e,
            #34495e 25%,
            #9b59b6 25%,
            #9b59b6 35%,
            #3498db 35%,
            #3498db 45%,
            #62cb31 45%,
            #62cb31 55%,
            #ffb606 55%,
            #ffb606 65%,
            #e67e22 65%,
            #e67e22 75%,
            #e74c3c 85%,
            #e74c3c 85%,
            #c0392b 85%,
            #c0392b 100%
        );
        background-size: 100% 6px;
        background-repeat: no-repeat;
        border-bottom: 1px solid #e4e5e7;
        overflow: auto;
        padding-bottom: 6px;
        text-align: center;
    }

    nav {
        border-bottom: 1px solid #e4e5e7;
        padding-top: 17px;
        padding-bottom: 15px;
        background: #f7f9fa;
        height: 60px;
        color: #6a6c6f;
        display: flex;
        justify-content: space-between;
    }

    nav a {
        margin-right: 1.5em;
        display: inline-block;
    }
</style>
