<template>
  <a-menu mode="vertical" @click="getNote">
    <a-sub-menu v-for="subMenu in subDir" :key="subMenu.key" :title="subMenu.key">
      <a-menu-item v-for="item in subMenu.files" :key="item.key">{{ item.key }}</a-menu-item>
    </a-sub-menu>
    <a-menu-item v-for="item in files" :key="item.key">{{ item.key }}</a-menu-item>
  </a-menu>
</template>

<script>
export default {
  name: "VMenu",
  data() {
    return {
      subDir: [],
      files: []
    }
  },
  props: {
    directory: '',
  },
  methods: {
    async updateMenu(dir) {
      this.subDir = [];
      let response = await this.axios.get('/' + dir + '/_vnote.json');
      if (response) {
        this.files = this.handleFiles(response.data.files);
        let subDirectories = response.data.sub_directories;
        for (let i = 0; i < subDirectories.length; i++) {
          let key = subDirectories[i].name;
          this.subDir.push({
            key,
            files: []
          });
        }
        for (let i = 0; i < this.subDir.length; i++) {
          let key = this.subDir[i].key;
          let subResponse = await this.axios.get('/' + dir + '/' + key + '/_vnote.json');
          if (subResponse) {
            this.subDir[i].files = this.handleFiles(subResponse.data.files);
          }
        }
      }
    },
    getNote(param) {
      let url = '/' + this.directory
      if (param.keyPath.length > 1) {
        url += '/' + param.keyPath[1];
      }
      url += '/' + param.keyPath[0] + '.md';
      document.title = param.keyPath[0];
      this.$emit('changeNote', url);
    },
    handleFiles(allFiles) {
      let files = [];
      for (let i = 0; i < allFiles.length; i++) {
        let name;
        if (allFiles[i].name.endsWith(".md")) {
          name = allFiles[i].name.substring(0, allFiles[i].name.length - 3);
        } else {
          name = allFiles[i].name;
        }
        files.push({
          key: name,
        });
      }
      return files;
    }
  },
  watch: {
    directory(newVal) {
      this.updateMenu(newVal);
    }
  },
  mounted() {
    while (this.directory === '') {
      console.log('empty');
    }
    this.updateMenu(this.directory);
  }
}
</script>

<style scoped>
.ant-menu {
  font-size: 16px;
  background-color: transparent;
}
</style>