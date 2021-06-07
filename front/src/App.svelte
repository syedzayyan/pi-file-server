<script>
	import router from 'page'
	import { auth } from "./components/auth";
    import { onDestroy } from "svelte";
	import Files from './routes/Files.svelte'
	import Layout from './routes/Layout.svelte'

	let auth_state;
    auth.subscribe((auth) => {
        auth_state = auth;
    });
    const unsubscribe = auth.subscribe((auth) => {
        auth_state = auth;
    });
	onDestroy(unsubscribe);
	// Include our Routes
  
	// Variables
	let page
	let params
	
	// Set up the pages to watch for
	router('/', () => {
		if (auth_state) {
			router.redirect("/drives/")
		} else {
			page = Files;
		}
	})
	router('/drives/*',
    (ctx, next) => {
      params = ctx.params
      next()
    },
    () => (page = Files))
	// Set up the router to start and actively watch for changes
	router.start()
</script>
<Layout>
  <svelte:component this="{page}" params="{params}" />
</Layout>  