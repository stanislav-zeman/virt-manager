<script>
    import {
        Sidebar,
        SidebarGroup,
        SidebarItem,
        SidebarWrapper,
    } from "flowbite-svelte";
    import { ChartPieSolid, UserSettingsSolid } from "flowbite-svelte-icons";
</script>

<Sidebar class="h-full">
    <SidebarWrapper class="h-full">
        <SidebarGroup>
            <SidebarItem label="Dashboard">
                <svelte:fragment slot="icon">
                    <ChartPieSolid
                        class="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                    />
                </svelte:fragment>
            </SidebarItem>
            <SidebarItem label="Settings">
                <svelte:fragment slot="icon">
                    <UserSettingsSolid
                        class="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                    />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>
    </SidebarWrapper>
</Sidebar>
