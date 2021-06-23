<template>
    <main role="main" class="container text-center">
        <h1>Create Account</h1>
        <notices :errors="errors" :notices="notices"></notices>
        <div class="row justify-content-md-center">
            <p>
            To keep things simple, we use a password-less approach.<br/>
            Every time you sign in, you will receive an email with link that will keep you signed in for some time.
            </p>
        </div>
        <div class="row justify-content-md-center">
            <form>
                <div class="form-row">
                    <div class="form-group">
                        <label for="form-fname" style="text-align: left">First Name</label>
                        <input v-model="input.fname" id="form-fname" class="form-control">
                    </div>
                </div>
                <div class="form-row">
                    <div class="form-group">
                        <label for="form-lname" style="text-align: left">Last Name</label>
                        <input v-model="input.lname" id="form-lname" class="form-control">
                    </div>
                </div>
                <div class="form-row">
                    <div class="form-group">
                        <label for="form-email" style="text-align: left">Email Address</label>
                        <input v-model="input.email" id="form-email" class="form-control">
                    </div>
                </div>
                <b-button type="button" v-on:click="register()" variant="primary">Register</b-button>
            </form>
        </div>
    </main>
</template>

<script>
import axios from 'axios';
import Notices from '../components/Notices.vue';
import AppConfig from '../util/config';
export default {
    components: { Notices },
    name: "register",
    data() {
        return {
            errors: [],
            notices: [],
            input: {
                fname: "",
                lname: "",
                email: ""
            }
        }
    },
    methods: {
        validateForm() {
            this.errors = [];
            this.input.fname = this.input.fname.trim()
            this.input.lname = this.input.lname.trim()
            this.input.email = this.input.email.trim()

            if (!this.input.fname.length) {
                this.errors.push("First name cannot be empty");
            }

            if (!this.input.lname.length) {
                this.errors.push("Last name cannot be empty");
            }

            var re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
            if (!re.test(this.input.email)) {
                this.errors.push("Email address is invalid");
            }

            if (!this.errors.length) { return true; }
            return false

        },
        register() {
            var successMessage = "Thank you for registering! You should receive welcome email and login instructions shortly.";
            var accountExistsMessage = "Oops, looks like this account already exists. Try signing in.";
            this.notices = [];
            if (!this.validateForm()) { window.scrollTo(0,0); return; }

            var config = {};
            axios.post(AppConfig.BACKEND_ENDPOINT + "/auth/signup", this.input, config).then(result => {
                var accountId = result.data.data;
                this.input = {};
                this.notices.push(successMessage);
                this.$emit("notice", successMessage);
                this.$emit("registered", accountId);
            }).catch(e => {
                if (e.response.data.error == "ACCOUNT_EXISTS") {
                    this.error.push(accountExistsMessage);
                } else {
                    this.errors.push(e.response.data.error);
                }
            }) 
        }
    },
    mounted() {
    }
}
</script>

<style scoped>
    #login {
        width: 500px;
        border: 1px solid #CCCCCC;
        background-color: #FFFFFF;
        margin: auto;
        margin-top: 200px;
        padding: 20px;
    }
</style>