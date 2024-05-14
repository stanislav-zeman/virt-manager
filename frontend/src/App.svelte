<script lang="ts">
    import Sidebar from "./layout/Sidebar.svelte";
    import Home from "./routes/Home.svelte";
    import Domains from "./routes/Domains.svelte";
    import Settings from "./routes/Settings.svelte";
    import Router from "svelte-spa-router";
    import NotFound from "./routes/NotFound.svelte";
    const routes = {
        "/": Home,
        "/domains": Domains,
        "/settings": Settings,
        "*": NotFound,
    };
</script>

<main class="w-full h-full flex">
    <Sidebar />
    <Router {routes} />
</main>
