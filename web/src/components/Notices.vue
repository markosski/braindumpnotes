<template>
    <div id="notices">
        <div class="errors" v-if="errors && errors.length">
            <div :key="error" v-for="error of errors" class="alert alert-warning alert-dismissible fade show" role="alert">
                {{error}}
            </div>
        </div>
        <div class="notifications" v-if="notices && notices.length">
            <div :key="notice" v-for="notice of notices" class="alert alert-success alert-dismissible fade show" role="alert">
                {{notice}}
            </div>
        </div>
    </div>
</template>
<script>
export default {
    name: 'notices',
    props: {
        errors: {
            type: Array,
            default: function() { return [] }
        },
        notices: {
            type: Array,
            default: function() { return [] }
        }
    },
    data() {
        return {}
    },
    methods: {}
}
</script>