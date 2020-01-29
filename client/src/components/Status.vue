<template>
    <div class="column container_sm">
        <div class="body_height">
            <div class="column">
                <h1 class="title">Server Status</h1>
            </div>
            <div class="column">
                <div class="column">
                    <div class="column">
                        Uptime: <strong>{{uptime}}</strong> seconds
                    </div>
                    <div class="column">
                        Session queries: <strong>{{sessionQueries}}</strong>
                    </div>
                </div>

                <div class="column">
                    <router-link to="/">
                        <button class="button">Back</button>
                    </router-link>

                </div>
            </div>

        </div>
    </div>

</template>

<script>

    export default {
        mounted() {
            this.loadServerStatus()

            setInterval(this.loadServerStatus, 1000)
        },
        data: () => ({
            uptime: "",
            sessionQueries: "",
            isLoading: false

        }),
        methods: {
            async loadServerStatus() {
                this.isLoading = true;
                try {
                    const resp = await this.$store.dispatch("serverStatus");
                    this.uptime = resp.data.uptime.toFixed(0);
                    this.sessionQueries = resp.data.sessionQueries;
                } catch (e) {
                    console.log(e);
                } finally {
                    this.isLoading = false;
                }
            }
        }
    }

</script>
