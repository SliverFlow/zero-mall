<template>
  <el-menu
    :default-active="defaultActive"
    :unique-opened="true"
    :collapse-transition="false"
    :collapse="asideStore.collapseType"
    router
    class="aside-con"
  >

    <template v-for="(v) in routerStore.menuList">
      <!--      拥有子菜单的-->
      <template v-if="v.children && v.children.length > 0">
        <el-sub-menu :index="v.path">
          <template #title>
            <el-icon>
              <component :is="v.icon" />
            </el-icon>
            <span>{{ v.title }}</span>
          </template>
          <template v-for="(iv) in v.children">
            <el-menu-item :index="iv.path">
              <el-icon>
                <component :is="iv.icon" />
              </el-icon>
              <span>{{ iv.title }}</span>
            </el-menu-item>
          </template>
        </el-sub-menu>
      </template>
      <!--无子菜单的-->
      <template v-if="v.children && v.children.length === 0">
        <el-tooltip
          :key="v.id"
          effect="dark"
          :content="v.title"
          placement="right"
          :disabled="!asideStore.collapseType"
        >
          <el-menu-item :key="v.id" :index="v.path">
            <el-icon>
              <component :is="v.icon" />
            </el-icon>
            <span>{{ v.title }}</span>
          </el-menu-item>
        </el-tooltip>
      </template>
    </template>

  </el-menu>
</template>

<script setup>
import { useASideStore } from '@/store/model/aside.js'
import { useRoute } from 'vue-router'
import { useRouterStore } from '@/store/model/router.js'
import { ref, watch } from 'vue'

const routerStore = useRouterStore()
const route = useRoute()
const asideStore = useASideStore()
const defaultActive = ref(route.path || '/layout/dashboard')

watch(route, () => {
  defaultActive.value = route.path
}, { immediate: true })
</script>

<style scoped lang="scss">
@import "@/style/variables.module";
@import "@/style/aside";

.aside-con {
  height: calc(100% - $TitleHeight);
}

</style>
