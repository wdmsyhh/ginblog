<template>
  <div>
    <Editor :init="init" v-model="content"></Editor>
  </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
import './tinymce.min.js'
import './icons/default/icons.min.js'
import '../../../public/tinymce/themes/silver/theme.min.js'

import './langs/zh_CN'

// 注册插件
import './plugins/preview/plugin.min.js'
import './plugins/paste/plugin.min.js'
import './plugins/wordcount/plugin.min.js'
import './plugins/code/plugin.min.js'

import './plugins/image/plugin.min.js'
import './plugins/imagetools/plugin.min.js'
import './plugins/media/plugin.min.js'
import './plugins/codesample/plugin.min.js'
import './plugins/lists/plugin.min.js'
import './plugins/table/plugin.min.js'

export default {
  components: { Editor },
  props: {
    value: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      init: {
        language: 'zh_CN',
        height: '600px',
        plugins: 'preview paste wordcount code imagetools image media codesample lists table',
        branding: false,
        paste_data_images: true,
        toolbar: [
          'undo redo | styleselect |fontsizeselect| bold italic underline strikethrough |alignleft aligncenter alignright alignjustify |blockquote removeformat |numlist bullist table',
          'preview paste code codesample |image media',
        ],
        //上传图片
        images_upload_handler: async (blobInfo, succFun, failFun) => {
          let formdata = new FormData()
          formdata.append('file', blobInfo.blob(), blobInfo.name())
          const { data: res } = await this.$http.post('upload', formdata)
          succFun(res.url)
          failFun(this.$message.error('上传图片失败'))
        },
        imagetools_cors_hosts: ['*'],
        imagetools_proxy: '',
        skin_url: '../../../tinymce/skins/ui/oxide', // skin路径，这里可以引用到 public 中的，虽然我也没有写前缀，疑惑,不过解决了打包部署后访问不到 css 文件的问题。
        content_css: '../../../tinymce/skins/content/default/content.min.css',
      },
      content: this.value,
    }
  },
  watch: {
    value(newV) {
      this.content = newV
    },
    content(newV) {
      this.$emit('input', newV)
    },
  },
}
</script>

<style scoped>
/* @import url('../../assets/tinymce/skins/ui/oxide/skin.min.css');
@import url('../../assets/tinymce/skins/content/default/content.min.css'); */
</style>
