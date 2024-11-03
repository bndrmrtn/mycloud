import type {Component} from "vue";
import { Mime } from 'mime/lite'
import standardTypes from 'mime/types/standard.js';
import otherTypes from 'mime/types/other.js';

// Components
// @ts-ignore
import FileIcon from "~/components/icons/file-icon.vue";
// @ts-ignore
import PythonFileType from "~/components/icons/filetypes/python-file-type.vue";
// @ts-ignore
import JsFileType from "~/components/icons/filetypes/js-file-type.vue";
// @ts-ignore
import SvgFileType from "~/components/icons/filetypes/svg-file-type.vue";
// @ts-ignore
import ZipFileType from "~/components/icons/filetypes/zip-file-type.vue";
// @ts-ignore
import ImageFileType from "~/components/icons/filetypes/image-file-type.vue";
// @ts-ignore
import JavaFileType from "~/components/icons/filetypes/java-file-type.vue";
// @ts-ignore
import GoFileType from "~/components/icons/filetypes/go-file-type.vue";

const mime = new Mime(standardTypes, otherTypes)
mime.define({ 'text/x-python': ['py'] })
mime.define({'text/x-go': ['go']})

export const getFileIcon = (type: string): Component => {
    if(type.includes(';')) type = type.split(';')[0]

    const ext = mime.getExtension(type)
    if(!ext) return FileIcon

    switch (ext) {
        case 'js':
            return JsFileType
        case 'svg':
            return SvgFileType
        case 'zip': case 'gz':
            return ZipFileType
        case 'py':
            return PythonFileType
        case 'java':
            return JavaFileType
        case 'go':
            return GoFileType
    }

    if(type.includes('image')) return ImageFileType

    return FileIcon
}