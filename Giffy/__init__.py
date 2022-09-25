from . import preferences

bl_info = {
    "name": "Giffy",
    "author": "Zivoy",
    "version": (1, 0, 0),
    "blender": (2, 82, 0),
    "location": "Properties > Render Properties",
    "description": "Render to a gif",
    "warning": "",
    "wiki_url": "",
    "tracker_url": "https://github.com/zivoy/giffy/issues",
    "support": "COMMUNITY",
    "category": "Render",
}



def register():
    preferences.register()


def unregister():
    preferences.unregister()


if __name__ == "__main__":
    register()
