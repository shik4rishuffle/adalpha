// Fix 'cannot find module ./App'
declare module "*.vue" {
    import Vue from "vue";
    export default Vue;
}