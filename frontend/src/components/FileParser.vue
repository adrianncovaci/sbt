<template>
  <div class="container">
    <b-card style="border:none" bg-variant="white" text-variant="dark">
      <b-button to="/export" router-tag="li">Export Policies</b-button>
      <b-button to="/manage" router-tag="li">Manage Policies</b-button>
      <hr class="my-4" />
      <b-card-text>Filename:</b-card-text>
      <b-form-input v-model="fileName" placeholder="Enter filename"></b-form-input>
      <hr class="my-4" />
      <b-card-text>Choose .audit:</b-card-text>
      <b-form-file
        :disabled="fileName === ''"
        v-model="file"
        ref="file-input"
        class="mb-2"
        @change="loadTextFromFile"
      ></b-form-file>
      <hr class="my-4" />
      <b-button :disabled="fileName === ''" variant="primary" @click="saveAudit">Save!</b-button>
      <hr class="my-4" />
    </b-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      file: null,
      fileName: ""
    };
  },
  methods: {
    loadTextFromFile(ev) {
      const file = ev.target.files[0];
      const reader = new FileReader();

      reader.onload = e => {
        var input = e.target.result;
        const fields = [
          "system",
          "type",
          "description",
          "cmd",
          "regex",
          "expect",
          "file",
          "info",
          "severity",
          "reference",
          "see_also",
          "severity",
          "solution"
        ];
        const xmlTags = /(<[^(<>)]*item>)([\s\S]*?)(<\/[^(<>)]*item>)/g;
        var output = input.match(xmlTags);
        for (var i = 0; i < output.length; i++) {
          output[i] = output[i].replace(/\n+/g, "\n");
          output[i] = output[i].split("\n");
          for (var j = 0; j < output[i].length; j++) {
            output[i][j] = output[i][j].replace(/ +(?= )/g, "");
            output[i][j] = output[i][j].replace(/(^ +)/g, "");
            var words = output[i][j].split(" ");
            if (!fields.includes(words[0].toLowerCase())) {
              output[i].splice(j, 1);
            }
          }
          output[i] = output[i].join("\n");
        }
        for (var k = 0; k < output.length; k++) {
          output[k] = output[k].split("\n");
          for (var m = 0; m < output[k].length; m++) {
            output[k][m] = output[k][m].replace(/(^ +)/g, "");
            var fwords = output[k][m].split(" ");
            if (!fields.includes(fwords[0].toLowerCase())) {
              output[k].splice(m, 1);
            }
          }
          output[k] = output[k].join("\n");
        }
        // AT THIS POINT, HAVE FILTERED ITEMS TO HAVE FIELDS WHICH CAN BE FOUND INSIDE "FIELDS" VARIABLE
        for (var l = 0; l < output.length; l++) {
          output[l] = output[l].split("\n");
          for (var p = 0; p < output[l].length; p++) {
            output[l][p] = output[l][p].replace(/(^ +)/g, "");
            output[l][p] = output[l][p].split(" ");
            var first_word = output[l][p][0];
            output[l][p][0] = `"${first_word}"`;
            output[l][p] = output[l][p].join(" ");
            output[l][p] = output[l][p].split(/:(.+)/);
            output[l][p][0] = output[l][p][0].trim();
            if (output[l][p][1]) {
              output[l][p][1] = output[l][p][1].replace(/"/g, "");
              output[l][p][1] = `"${output[l][p][1]}",`;
            }
            output[l][p].pop();
            output[l][p] = output[l][p].join(":");
          }
          output[l] = output[l].join("\n");
          output[l].replace(/^\s*\n/gm, "");
          output[l].replace("\\", "\\\\");
          output[l] = output[l].substring(0, output[l].length - 1);
        }
        // CREATING JSON OBJECTS
        var items = [];
        output.forEach(element => {
          var properties = element.split(",\n");
          var obj = {};
          properties.forEach(function(property) {
            var tup = property.split(":");
            tup[0] = tup[0].replace(/"/g, "");
            tup[1] = tup[1].replace(/"/g, "");
            obj[tup[0]] = tup[1];
          });
          items.push(obj);
        });
        this.saveFile(items);
      };
      reader.readAsText(file);
    },
    saveFile: function(items) {
      const data = JSON.stringify(items);
      const blob = new Blob([data], { type: "text/plain" });
      const e = document.createEvent("MouseEvents"),
        a = document.createElement("a");
      a.download = `${this.fileName}.json`;
      a.href = window.URL.createObjectURL(blob);
      a.dataset.downloadurl = ["text/json", a.download, a.href].join(":");
      e.initEvent(
        "click",
        true,
        false,
        window,
        0,
        0,
        0,
        0,
        0,
        false,
        false,
        false,
        false,
        0,
        null
      );
      a.dispatchEvent(e);
    },
    saveAudit() {
      const filename = `${this.fileName}.json`;
      window.backend.basic(filename).then(console.log("sent:"));
    }
  }
};
</script>

<style>
</style>