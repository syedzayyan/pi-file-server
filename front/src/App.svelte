<script>
	import router from 'page'
  
	// Include our Routes
	import Files from './routes/Files.svelte'
import Layout from './routes/Layout.svelte'
  
	// Variables
	let page
	let params
  
	// Set up the pages to watch for
	router('/', () => (page = Files))
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