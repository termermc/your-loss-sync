package lang

// DefaultLangCode is the default language code.
// This will be used if no translation is found for another language.
// All translations must support this language.
const DefaultLangCode = "en-us"

// Languages is a mapping of supported language codes to their names.
var Languages = map[string]string{
	"en-us":  "English (US)",
	"es-419": "Español (Latinoamérica)",
	"zh-cn":  "中文 (中国)",
}

// GetLangNames returns a list of all supported language names.
func GetLangNames() []string {
	names := make([]string, 0, len(Languages))
	for _, name := range Languages {
		names = append(names, name)
	}
	return names
}

// GetLangCodeFromName returns the language code for the given language name.
// If the name is not found, an empty string is returned.
func GetLangCodeFromName(name string) string {
	for code, n := range Languages {
		if n == name {
			return code
		}
	}

	return ""
}

// Translations is a map of translation keys to a map of language codes to translations.
// Translation strings can have parameters, which are denoted by $1, $2, etc.
// The same parameters can be repeated multiple times or not used at all.
var Translations = map[string]map[string]string{
	"general.confirm": {
		"en-us":  "Confirm",
		"es-419": "Confirmar",
		"zh-cn":  "确认",
	},
	"general.cancel": {
		"en-us":  "Cancel",
		"es-419": "Cancelar",
		"zh-cn":  "取消",
	},
	"general.create": {
		"en-us":  "Create",
		"es-419": "Crear",
		"zh-cn":  "创建",
	},
	"general.save": {
		"en-us":  "Save",
		"es-419": "Guardar",
		"zh-cn":  "保存",
	},
	"general.error": {
		"en-us":  "Error",
		"es-419": "Error",
		"zh-cn":  "错误",
	},

	"widget.file-picker.select-file": {
		"en-us":  "Select File",
		"es-419": "Seleccionar Archivo",
		"zh-cn":  "选择文件",
	},
	"widget.file-picker.select-folder": {
		"en-us":  "Select Folder",
		"es-419": "Seleccionar Carpeta",
		"zh-cn":  "选择文件夹",
	},

	"profile.default.hq-mp3.name": {
		"en-us":  "High-Quality MP3",
		"es-419": "MP3 de Alta Calidad",
		"zh-cn":  "高质量 MP3",
	},
	"profile.default.flac.name": {
		"en-us":  "Lossless FLAC",
		"es-419": "FLAC sin pérdidas",
		"zh-cn":  "无损 FLAC",
	},
	"profile.default.wav.name": {
		"en-us":  "Lossless WAV",
		"es-419": "WAV sin pérdidas",
		"zh-cn":  "无损 WAV",
	},
	"profile.default.hq-aac.name": {
		"en-us":  "High-Quality AAC",
		"es-419": "AAC de Alta Calidad",
		"zh-cn":  "高质量 AAC",
	},
	"profile.default.alac.name": {
		"en-us":  "Lossless ALAC",
		"es-419": "ALAC sin pérdidas",
		"zh-cn":  "无损 ALAC",
	},

	"shell.tab.syncs": {
		"en-us":  "Syncs",
		"es-419": "Sincronizaciones",
		"zh-cn":  "同步",
	},
	"shell.tab.profiles": {
		"en-us":  "Profiles",
		"es-419": "Perfiles",
		"zh-cn":  "配置文件",
	},
	"shell.tab.settings": {
		"en-us":  "Settings",
		"es-419": "Ajustes",
		"zh-cn":  "设置",
	},
	"shell.tab.in-progress": {
		"en-us":  "In Progress ($1/$2)",
		"es-419": "En progreso ($1/$2)",
		"zh-cn":  "进行中 ($1/$2)",
	},

	"config.error.unsupported-version": {
		"en-us":  "Unsupported config version",
		"es-419": "Versión de configuración no soportada",
		"zh-cn":  "不支持的配置版本",
	},
	"config.error.unknown-profile": {
		"en-us":  "Config contains a reference to an unknown profile",
		"es-419": "Configuración contiene una referencia a un perfil desconocido",
		"zh-cn":  "配置包含对未知配置文件的引用",
	},

	"setup.title": {
		"en-us":  "Setup",
		"es-419": "Configuración",
		"zh-cn":  "设置",
	},
	"setup.select-language": {
		"en-us":  "Select Language",
		"es-419": "Seleccionar Idioma",
		"zh-cn":  "选择语言",
	},

	"tab.syncs.create": {
		"en-us":  "Create Sync",
		"es-419": "Crear Sincronización",
		"zh-cn":  "创建同步",
	},
	"tab.syncs.delete-confirm.title": {
		"en-us":  "Delete Sync",
		"es-419": "Eliminar Sincronización",
		"zh-cn":  "删除同步",
	},
	"tab.syncs.delete-confirm.description": {
		"en-us":  "Are you sure you want to delete the sync \"$1\"?",
		"es-419": "¿Estás seguro de que quieres eliminar la sincronización \"$1\"?",
		"zh-cn":  "您确定要删除同步 \"$1\"？",
	},
	"tab.syncs.form.name": {
		"en-us":  "Name",
		"es-419": "Nombre",
		"zh-cn":  "名称",
	},
	"tab.syncs.form.source-dir": {
		"en-us":  "Source Directory",
		"es-419": "Directorio de Origen",
		"zh-cn":  "源目录",
	},
	"tab.syncs.form.dest-dir": {
		"en-us":  "Destination Directory",
		"es-419": "Directorio de Destino",
		"zh-cn":  "目标目录",
	},
	"tab.syncs.form.profile": {
		"en-us":  "Profile",
		"es-419": "Perfil",
		"zh-cn":  "配置文件",
	},
	"tab.syncs.form.escape-filenames": {
		"en-us":  "Replace invalid characters in filenames?",
		"es-419": "¿Reemplazar caracteres no válidos en los nombres de archivos?",
		"zh-cn":  "替换文件名中的无效字符？",
	},
	"tab.syncs.form.reencode-same-format": {
		"en-us":  "Reencode files with the same format?",
		"es-419": "¿Reencodificar archivos con el mismo formato?",
		"zh-cn":  "重新编码具有相同格式的文件？",
	},
	"tab.syncs.form.error.missing-name": {
		"en-us":  "Name is required",
		"es-419": "Se requiere el nombre",
		"zh-cn":  "名称为必填项",
	},
	"tab.syncs.form.error.missing-source-dir": {
		"en-us":  "Source directory is required",
		"es-419": "Se requiere el directorio de origen",
		"zh-cn":  "源目录为必填项",
	},
	"tab.syncs.form.error.missing-dest-dir": {
		"en-us":  "Destination directory is required",
		"es-419": "Se requiere el directorio de destino",
		"zh-cn":  "目标目录为必填项",
	},
	"tab.syncs.form.error.missing-profile": {
		"en-us":  "Profile is required. Go create a profile first.",
		"es-419": "Se requiere el perfil. Primero cree un perfil.",
		"zh-cn":  "配置文件为必填项。首先创建配置文件。",
	},
	"tab.syncs.form.error.name-exists": {
		"en-us":  "A sync with the same name already exists",
		"es-419": "Ya existe una sincronización con el mismo nombre",
		"zh-cn":  "同名同步已存在",
	},
	"tab.syncs.form.error.source-dir-nonexistent": {
		"en-us":  "Source directory does not exist, or is not a directory",
		"es-419": "El directorio de origen no existe o no es un directorio",
		"zh-cn":  "源目录不存在或不是目录",
	},
	"tab.syncs.form.error.dest-dir-nonexistent": {
		"en-us":  "Destination directory does not exist, or is not a directory",
		"es-419": "El directorio de destino no existe o no es un directorio",
		"zh-cn":  "目标目录不存在或不是目录",
	},
	"tab.syncs.form.error.source-dest-dirs-same": {
		"en-us":  "Source and destination directories cannot be the same",
		"es-419": "Los directorios de origen y destino no pueden ser los mismos",
		"zh-cn":  "源和目标目录不能相同",
	},

	"tab.profiles.create": {
		"en-us":  "Create Profile",
		"es-419": "Crear Perfil",
		"zh-cn":  "创建配置文件",
	},
	"tab.profiles.form.name": {
		"en-us":  "Name",
		"es-419": "Nombre",
		"zh-cn":  "名称",
	},
	"tab.profiles.form.format": {
		"en-us":  "Format",
		"es-419": "Formato",
		"zh-cn":  "格式",
	},
	"tab.profiles.form.is-lossless": {
		"en-us":  "Lossless",
		"es-419": "Sin pérdidas",
		"zh-cn":  "无损",
	},
	"tab.profiles.form.supports-metadata": {
		"en-us":  "Supports Metadata",
		"es-419": "Soporta metadatos",
		"zh-cn":  "支持元数据",
	},
	"tab.profiles.form.supports-artwork": {
		"en-us":  "Supports Artwork",
		"es-419": "Soporta arte",
		"zh-cn":  "支持艺术",
	},
	"tab.profiles.form.bitrate": {
		"en-us":  "Bitrate",
		"es-419": "Bitrate",
		"zh-cn":  "比特率",
	},
	"tab.profiles.form.error.invalid-bitrate": {
		"en-us":  "Invalid bitrate",
		"es-419": "Bitrate inválido",
		"zh-cn":  "无效的比特率",
	},
	"tab.profiles.form.error.missing-name": {
		"en-us":  "Name is required",
		"es-419": "Se requiere el nombre",
		"zh-cn":  "名称为必填项",
	},
	"tab.profiles.form.error.name-exists": {
		"en-us":  "A profile with the same name already exists",
		"es-419": "Ya existe un perfil con el mismo nombre",
		"zh-cn":  "同名配置文件已存在",
	},
	"tab.profiles.delete-confirm.title": {
		"en-us":  "Delete Profile",
		"es-419": "Eliminar Perfil",
		"zh-cn":  "删除配置文件",
	},
	"tab.profiles.delete-confirm.description": {
		"en-us":  "Are you sure you want to delete the profile \"$1\"?",
		"es-419": "¿Estás seguro de que quieres eliminar el perfil \"$1\"?",
		"zh-cn":  "您确定要删除配置文件 \"$1\"？",
	},
	"tab.profiles.error.in-use": {
		"en-us":  "Profile is in use by a sync. Delete syncs using this profile first.",
		"es-419": "El perfil está en uso por una sincronización. Elimina primero las sincronizaciones que usan este perfil.",
		"zh-cn":  "配置文件正在被同步使用。先删除同步使用的配置文件。",
	},

	"tab.progress.start": {
		"en-us":  "Start",
		"es-419": "Iniciar",
		"zh-cn":  "开始",
	},
	"tab.progress.cancel": {
		"en-us":  "Cancel",
		"es-419": "Cancelar",
		"zh-cn":  "取消",
	},
	"tab.progress.status-label": {
		"en-us":  "Completed: $1/$2, failed: $3",
		"es-419": "Completado: $1/$2, fallidos: $3",
		"zh-cn":  "已完成: $1/$2, 失败: $3",
	},

	"sync.scanning-source": {
		"en-us":  "Scanning source directory...",
		"es-419": "Escaneando directorio de origen...",
		"zh-cn":  "正在扫描源目录...",
	},
	"sync.copying": {
		"en-us":  "Copying $1",
		"es-419": "Copiando $1",
		"zh-cn":  "正在复制 $1",
	},
	"sync.transcoding": {
		"en-us":  "Transcoding $1",
		"es-419": "Transcodificando $1",
		"zh-cn":  "正在转码 $1",
	},
	"sync.path-already-exists": {
		"en-us":  "Path $1 already exists, skipping",
		"es-419": "La ruta $1 ya existe, se omite",
		"zh-cn":  "路径 $1 已存在，将跳过",
	},
	"sync.done": {
		"en-us":  "Done (total: $1, completed: $2, failed: $3)",
		"es-419": "Hecho (total: $1, completados: $2, fallidos: $3)",
		"zh-cn":  "完成 (总计: $1, 已完成: $2, 失败: $3)",
	},
}
