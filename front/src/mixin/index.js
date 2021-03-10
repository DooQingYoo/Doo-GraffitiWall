const mixin = {
    methods: {
        async axiosGet(url) {
            try {
                const response = await this.axios.get(url);
                if (response.data) {
                    return response.data;
                } else {
                    return null;
                }
            } catch (e) {
                this.$message.error("网络错误，请稍后重试")
                return null;
            }
        }
    },
}

export default mixin;