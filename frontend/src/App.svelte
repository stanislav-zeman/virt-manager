<script lang="ts">
    import Sidebar from "./components/sidebar.svelte";
    import Domains from "./pages/domains.svelte";
    import Settings from "./pages/settings.svelte";
    import Router from "svelte-spa-router";
    const routes = {
        "/domains": Domains,
        "/settings": Settings,
    };
</script>

<main class="w-full h-full flex">
    <Sidebar />
    <Router {routes} />
</main>
