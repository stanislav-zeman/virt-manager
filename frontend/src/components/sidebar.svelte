<script lang="ts">
    import {
        Sidebar,
        SidebarGroup,
        SidebarItem,
        SidebarWrapper,
    } from "flowbite-svelte";
    import { DesktopPcSolid, UserSettingsSolid } from "flowbite-svelte-icons";
</script>

<Sidebar class="w-1/5">
    <SidebarWrapper class="h-full">
        <SidebarGroup>
            <SidebarItem label="Domains" href="#/domains">
                <svelte:fragment slot="icon">
                    <DesktopPcSolid
                        class="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                    />
                </svelte:fragment>
            </SidebarItem>
            <SidebarItem label="Settings" href="#/settings">
                <svelte:fragment slot="icon">
                    <UserSettingsSolid
                        class="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                    />
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>
    </SidebarWrapper>
</Sidebar>
