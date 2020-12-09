// Code generated by file2byteslice. DO NOT EDIT.

package main

var water_go = []byte("// Copyright 2020 The Ebiten Authors\n//\n// Licensed under the Apache License, Version 2.0 (the \"License\");\n// you may not use this file except in compliance with the License.\n// You may obtain a copy of the License at\n//\n//     http://www.apache.org/licenses/LICENSE-2.0\n//\n// Unless required by applicable law or agreed to in writing, software\n// distributed under the License is distributed on an \"AS IS\" BASIS,\n// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n// See the License for the specific language governing permissions and\n// limitations under the License.\n\n// +build ignore\n\npackage main\n\nvar Time float\nvar Cursor vec2\nvar ScreenSize vec2\n\nfunc Fragment(position vec4, texCoord vec2, color vec4) vec4 {\n\tborder := ScreenSize.y*0.6 + 4*cos(Time*3+texCoord.y*200)\n\tif position.y < border {\n\t\treturn imageSrc2UnsafeAt(texCoord)\n\t}\n\n\tsrcsize := imageSrcTextureSize()\n\trorigin, _ := imageSrcRegionOnTexture()\n\n\txoffset := (4 / srcsize.x) * cos(Time*3+texCoord.y*200)\n\tyoffset := (20 / srcsize.y) * (1.0 + cos(Time*3+texCoord.y*50))\n\tbordertex := border / srcsize.y\n\tclr := imageSrc2At(vec2(\n\t\ttexCoord.x+xoffset,\n\t\t-(texCoord.y+yoffset-rorigin.y)+bordertex*2+rorigin.y,\n\t)).rgb\n\n\toverlay := vec3(0.5, 1, 1)\n\treturn vec4(mix(clr, overlay, 0.25), 1)\n}\n")
