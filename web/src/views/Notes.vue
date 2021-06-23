<template>
    <main role="main" class="container">
        <h1>Your Notes</h1>
        <notices :errors="errors" :notices="notices"></notices>
        <div class="row justify-content-md-center">
        </div>
        <form class="form-inline">
            <input v-model="newNote.name" id="form-note" class="form-control m-2" placeholder="Note File Name">
            <b-button type="button" class="btn btn-primary m-2" v-on:click="makeNote()">New Note</b-button><br/>
        </form>
        <div>
            <div v-for="note in notes" :key="note.name">
                <button v-on:click="editNote(note.name)">{{ note.name }}</button>
            </div>
        </div>
        <div id="editor" tabindex="0" @keydown.ctrl.83.prevent.stop="doSave">
            <textarea v-if="current" :value="current.content" @input="update"></textarea>
            <div class="marked" v-html="compiledMarkdown"></div>
        </div>
        <div class="row justify-content-md-center">
        </div>
    </main>
</template>

<script>
const marked = require('marked');
import _ from 'lodash';
import Notices from '../components/Notices.vue';
import CommonMixin from '../util/common';
// import AppConfig from '../util/config';
export default {
    components: { Notices },
    name: "notes",
    mixins: [CommonMixin],
    data() {
        return {
            errors: [],
            notices: [],
            notes: [],
            fakeNotes: [
                {
                    name: "foo.md",
                    content: "# Foo\n\nHello World"
                },
                {
                    name: "bar.md",
                    content: "# Bar"
                }
            ],
            current: null,
            newNote: {
                name: "",
                content: "",
                format: "MD"
            }
        }
    },
    computed: {
        compiledMarkdown() {
            if (!this.current) {
                return marked("");
            } else {
                return marked(this.current.content);
            }
        }
    },
    methods: {
        makeNote() {
            this.notes.push(this.newNote);
            this.current = this.newNote;
            this.newNote = {
                name: "",
                content: ""
            };

            var path = "/notes";
            this.backendPost(path, this.current).then(() => {
            })
        },
        create() {
            var path = "/notes";
            this.backendPost(path, this.current).then(() => {
                // console.log(result);
            });
        },
        doSave(e) {
            if (!(e.keyCode === 83 && (e.ctrlKey || e.metaKey))) {
                return;
            }
            var path = "/notes/" + this.current.name;
            var update = {
                name: this.current.name, 
                content: this.current.content
            }
            this.backendPatch(path, update).then(() => {
            });

            e.preventDefault();
        },
        delete() {
            // Delete it
        },
        load: function() {
            var that = this;
            this.backendGet("/notes?accountId=123").then(result => {
                that.notes = result.data;
            });
        },
        update: _.debounce(function(e) {
            this.current.content = e.target.value;
        }, 300),
        editNote: function(noteName) {
            const note = _.find(this.notes, function(o) { return o.name == noteName});
            this.current = note;
        }
    },
    beforeDestroy() {
        document.removeEventListener("keydown", this.doSave);
    },
    mounted() {
        document.addEventListener("keydown", this.doSave);
        this.load();
    }
}
</script>

<style scoped>
    html,
    body,
    #editor {
    margin: 0;
    height: 400px;
    font-family: "Helvetica Neue", Arial, sans-serif;
    color: #333;
    }

    textarea,
    #editor div {
    display: inline-block;
    width: 49%;
    height: 100%;
    vertical-align: top;
    box-sizing: border-box;
    padding: 0 20px;
    }

    textarea {
    border: none;
    border-right: 1px solid #ccc;
    resize: none;
    outline: none;
    background-color: #f6f6f6;
    font-size: 14px;
    font-family: "Monaco", courier, monospace;
    padding: 20px;
    }

    .marked {
        zoom: 1;
    }

    code {
    color: #f66;
    }
</style>