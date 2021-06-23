<template>
    <main role="main" class="container text-center">
        <h1>Sign In</h1>
        <notices :errors="errors" :notices="notices"></notices>
        <div class="row justify-content-md-center">
            <p>Sign in and check your mailbox for access link.
            </p>
        </div>
        <div class="row justify-content-md-center">
            <form class="form-inline col-md-4">
                <!-- <div class="form-group"> -->
                    <input v-model="input.email" id="form-email" class="form-control col-md-9" placeholder="Email">
                    <b-button type="button" class="btn btn-primary" v-on:click="emailLogin()" variant="primary" :disabled="formCompleted">Sign In</b-button>
                <!-- </div> -->
            </form>
        </div>
        <div class="row justify-content-md-center">
            <div>or <router-link to="/register">Create Account</router-link></div>
        </div>
    </main>
</template>

<script>
import axios from 'axios';
import Notices from '../components/Notices.vue';
import AppConfig from '../util/config';
export default {
    components: { Notices },
    name: 'Login',
    data() {
        return {
            errors: [],
            notices: [],
            input: {
                email: "",
            },
            formCompleted: false
        }
    },
    methods: {
        login() {
            this.errors = [];

            if (this.$route.query.token && this.$route.query.email) {
                var self = this;
                var token = this.$route.query.token;
                var email = this.$route.query.email;
                var config = {
                    params: {
                        "token": token,
                        "email": email
                    }
                };
                axios.get(AppConfig.BACKEND_ENDPOINT + "/auth/authorize", config)
                .then(result => {
                    document.cookie = "token=" + token;
                    self.$emit("authenticated", true);
                    self.$router.replace({ name: "dashboard" });
                    self.formCompleted = true;
                }).catch(e => {
                    console.log(e);
                    if (e.response.status === 403) {
                        this.$emit("authenticated", false);
                        this.$router.replace({ name: "home" });
                    } else if (e.response.status == 503) {
                        self.errors.push("There was an internal error. Please try again in a while.");
                    } else {
                        self.errors.push(e.response.data.error);
                    }
                }) 
            }
        },
        validateForm() {
            this.errors = [];
            this.input.email = this.input.email.trim()

            var re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
            if (!re.test(this.input.email)) {
                this.errors.push("Email address is invalid");
            }

            if (!this.errors.length) { return true; }
            return false

        },
        emailLogin() {
            this.notices = [];
            this.errors = [];
            if (!this.validateForm()) { window.scrollTo(0,0); return; }

            var config = {
                params: {
                    "email": this.input.email
                }
            };
            axios.get(AppConfig.BACKEND_ENDPOINT + "/auth/login", config).then(result => {
                this.input = {};
                this.notices.push("Access link has been sent to your email address.");
            }).catch(e => {
                this.errors.push(e.response.data.error);
            }) 
        }
    },
    mounted() {
        this.login();
    }
}
</script>

<style scoped>
</style>