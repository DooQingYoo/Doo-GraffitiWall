<template>
  <div>
    <a-affix :offset-top="0">
      <a-menu mode="horizontal" @click="goTo">
        <a-menu-item key="book" style="border-bottom: 2px solid transparent;">
          <img src="@/assets/image/home.png"/>
        </a-menu-item>
        <a-menu-item v-for="item in items" :key="item.name">{{ item.name }}</a-menu-item>
      </a-menu>
    </a-affix>
  </div>
</template>

<script>

export default {
  name: "Menu",
  data() {
    return {
      items: []
    }
  },
  methods: {
    async getMenuItems() {
      this.items = await this.axiosGet("/");
    },
    goTo(param) {
      if ('book' === param.key) {
        this.$router.replace('/');
      } else {
        this.$router.push('/showNote/' + param.key);
      }

    }
  },
  created() {
    this.getMenuItems();
  }
}
</script>

<style scoped>
.ant-menu {
  font-size: 18px;
}

.ant-menu img {
  height: 30px;
  width: 30px;
}
</style>