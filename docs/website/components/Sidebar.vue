<template>
  <div class="sidenav sticky pt-4 pb-4">
    <ul class="mt-8">
      <li
        v-for="entry in $store.state.sidebar.menu"
        :key="entry.title"
        @click="selected = entry.title"
        class="py-2"
      >
        <a
          :href="'#' + entry.path"
          @click="handleClick(entry)"
          class="sidebar-category"
        >
          <span class="relative">{{ entry.title }}</span>
        </a>
        <ul class="py-0 pl-4">
          <li v-for="item in entry.items" :key="item.path" class="ml-0">
            <a
              :href="'#' + item.path"
              @click="handleClick(item)"
              class="sidebar-item"
            >
              <span class="relative">{{ item.title }}</span>
            </a>
            <ul class="py-0 pl-4">
              <li v-for="child in item.children" :key="child.path" class="ml-0">
                <a
                  :href="'#' + child.path"
                  @click="handleClick(child)"
                  class="sidebar-child"
                >
                  <span class="relative">{{ child.title }}</span>
                </a>
              </li>
            </ul>
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Sidebar',

  data() {
    return {
      selected: undefined
    }
  },

  methods: {
    handleClick(item) {
      this.$store.commit('sidebar/setActiveDocPath', item.path)
    }
  }
}
</script>

<style>
.sidenav {
  top: 7%;
  overflow-x: hidden;
}

a.sidebar-category {
  @apply font-brand text-gray-600 relative block text-lg font-bold tracking-wide;
}

a.sidebar-category:hover {
  @apply text-gray-900;
}

a.sidebar-item {
  @apply font-brand text-gray-600 relative block text-base tracking-wide;
}

a.sidebar-item:hover {
  @apply text-gray-900;
}

a.sidebar-child {
  @apply font-brand text-gray-600 relative block text-xs tracking-wide;
}

a.sidebar-child:hover {
  @apply text-gray-900;
}

a.active {
  @apply text-gray-900 font-bold;
}
</style>
