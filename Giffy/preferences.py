import os.path
import shutil
import subprocess

import bpy

addon_idname = __package__.split(".")[0]


def getPreferences(context=None):
    if context is None:
        context = bpy.context
    preferences = context.preferences
    addon_preferences = preferences.addons[addon_idname].preferences
    return addon_preferences


tmp = shutil.which("ffmpeg")
ffmpeg_path_default = tmp if tmp is not None else ""
tmp = shutil.which("gifsicle")
gifsicle_path_default = tmp if tmp is not None else ""


def isValid(itemName, args, lookfor):
    def valid(self):
        path = getattr(self, itemName)
        return validFile(path, args, lookfor)

    return valid


ffmpeg_lookfor_args = "-version"
ffmpeg_lookfor_value = "ffmpeg"
gifsicle_lookfor_args = "--version"
gifsicle_lookfor_value = "Gifsicle"


class GiffyPreferences(bpy.types.AddonPreferences):
    bl_idname = addon_idname

    def ffempegSet(self, context):
        validFile(self.ffmpeg_path, ffmpeg_lookfor_args, ffmpeg_lookfor_value)

    ffmpeg_path: bpy.props.StringProperty(
        name="FFmpeg Path",
        subtype="FILE_PATH",
        default=ffmpeg_path_default,
        update=ffempegSet,
    )

    ffmpeg_valid: bpy.props.BoolProperty(
        name="Valid FFmpeg Path",
        get=isValid("ffmpeg_path", ffmpeg_lookfor_args, ffmpeg_lookfor_value)
    )

    def gifsicleSet(self, context):
        validFile(self.gifsicle_path, gifsicle_lookfor_args, gifsicle_lookfor_value)

    gifsicle_path: bpy.props.StringProperty(
        name="Gifsicle Path",
        subtype="FILE_PATH",
        default=gifsicle_path_default,
        update=ffempegSet,
    )

    gifsicle_valid: bpy.props.BoolProperty(
        name="Valid Gifsicle Path",
        get=isValid("gifsicle_path", gifsicle_lookfor_args, gifsicle_lookfor_value)
    )

    def draw(self, context):
        layout = self.layout
        layout.label(text="FFmpeg is required, it is used to construct the gif")
        layout.prop(self, "ffmpeg_path")
        layout.operator("wm.url_open", text="Download FFmpeg").url = "https://ffmpeg.org/download.html"

        layout.label(text="Gifsicle is used to optimise the gif")
        layout.prop(self, "gifsicle_path")
        layout.operator("wm.url_open", text="Download Gifsicle").url = "https://www.lcdf.org/gifsicle/"


# prevent rechecking
paths = dict()


def validFile(path, args, lookfor):
    if not os.path.isfile(path):
        return False

    if path in paths:
        return paths[path]

    res = subprocess.Popen([path, args], shell=False, stdout=subprocess.PIPE).stdout.read().decode()
    paths[path] = lookfor in res
    return paths[path]


register, unregister = bpy.utils.register_classes_factory((GiffyPreferences,))
