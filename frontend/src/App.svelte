<script lang="ts">
    import Sidebar from "./layout/Sidebar.svelte";
    import Home from "./routes/Home.svelte";
    import Settings from "./routes/Settings.svelte";
    import Router from "svelte-spa-router";
    import { wrap } from "svelte-spa-router/wrap";
    import NotFound from "./routes/NotFound.svelte";
    const routes = {
        "/": Home,
        "/domains": wrap({
            asyncComponent: () => import("./routes/Domains.svelte"),
        }),
        "/settings": Settings,
        "*": NotFound,
    };
</script>

<main class="w-full h-full flex">
    <Sidebar />
    <Router {routes} />
</main>
