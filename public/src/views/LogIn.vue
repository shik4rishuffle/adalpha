<template>
    <div class="container">
        <div class="row justify-content-center align-items-center vh-100">
            <div class="col-sm-11 col-md-6 card">
                <h1 class="h1 align-self-center">Welcome</h1>
                <form class="col-12" @submit.prevent="login">
                    <div class="form-group">
                        <label >Username</label>
                        <input v-model="username" class="form-control align-self-center" type="text"/>
                    </div>

                    <div class="form-group">
                        <label>Password</label>
                        <input v-model="password" class="form-control align-self-center" type="password"/>
                    </div>
                    <button type="submit" class="btn btn-primary mb-3">Submit</button>
                    <span class="float-right text-danger">{{this.loginError}}</span>
                </form>
            </div>
            <div>
        </div>
    </div>
    </div>
</template>

<script lang="ts">
  import {Vue, Component} from 'vue-property-decorator';

  @Component
  export default class LogIn extends Vue {
        username = '';
        password = '';
        loginError = '';
    // methods
        login () {
            fetch('http://localhost:8080/login', {
                method: 'POST',
                body: JSON.stringify({username: this.username, password: this.password}),
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: 'include'
            }).then(res => {
                if (res.status === 200) {
                    return 'Success';
                } else if (res.status === 401) {
                    throw new Error('username or password incorrect');
                }
            })
                .then(() => {
                    this.$store.dispatch('logInHandler', this.username);
                    this.$router.push('/DashBoard');
                })
                .catch((error) => {
                       this.loginError = error;
                    console.log(this.loginError);
                });
        }
  };
</script>

<style scoped>
.card {
    border:none;
}
@media only screen and (min-width: 768px) {
    .card {
        border: 1px solid #ccc;
    }
}
</style>
