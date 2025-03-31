package lang

// DefaultLangCode is the default language code.
// This will be used if no translation is found for another language.
// All translations must support this language.
const DefaultLangCode = "en-us"

// Languages is a mapping of supported language codes to their names.
var Languages = map[string]string{
	"en-us":  "English (US)",
	"es-419": "Español (Latinoamérica)",
	"fr-fr":  "Francais (France)",
	"it-it":  "Italiano (Italia)",
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
		"fr-fr":  "Confirmer",
		"it-it":  "Confermare",
		"zh-cn":  "确认",
	},
	"general.cancel": {
		"en-us":  "Cancel",
		"es-419": "Cancelar",
		"fr-fr":  "Annuler",
		"it-it":  "Annullare",
		"zh-cn":  "取消",
	},
	"general.create": {
		"en-us":  "Create",
		"es-419": "Crear",
		"fr-fr":  "Creer",
		"it-it":  "Creare",
		"zh-cn":  "创建",
	},
	"general.save": {
		"en-us":  "Save",
		"es-419": "Guardar",
		"fr-fr":  "Sauvegarder",
		"it-it":  "Salvare",
		"zh-cn":  "保存",
	},
	"general.error": {
		"en-us":  "Error",
		"es-419": "Error",
		"fr-fr":  "Erreur",
		"it-it":  "Errore",
		"zh-cn":  "错误",
	},
	"general.warning": {
		"en-us":  "Warning",
		"es-419": "Advertencia",
		"fr-fr":  "Avertissement",
		"it-it":  "Avviso",
		"zh-cn":  "警告",
	},

	"widget.file-picker.select-file": {
		"en-us":  "Select File",
		"es-419": "Seleccionar Archivo",
		"fr-fr":  "Selectionner fichier",
		"it-it":  "Selezionare documento",
		"zh-cn":  "选择文件",
	},
	"widget.file-picker.select-folder": {
		"en-us":  "Select Folder",
		"es-419": "Seleccionar Carpeta",
		"fr-fr":  "Selectionner dossier",
		"it-it":  "Selezionare cartella",
		"zh-cn":  "选择文件夹",
	},

	"profile.default.hq-mp3.name": {
		"en-us":  "High-Quality MP3",
		"es-419": "MP3 de Alta Calidad",
		"fr-fr":  "MP3 de haute qualite",
		"it-it":  "MP3 di alta qualita'",
		"zh-cn":  "高质量 MP3",
	},
	"profile.default.flac.name": {
		"en-us":  "Lossless FLAC",
		"es-419": "FLAC sin pérdidas",
		"fr-fr":  "FLAC sans perte",
		"it-it":  "FLAC senza perdita",
		"zh-cn":  "无损 FLAC",
	},
	"profile.default.wav.name": {
		"en-us":  "Lossless WAV",
		"es-419": "WAV sin pérdidas",
		"fr-fr":  "WAV sans perte",
		"it-it":  "WAV senza perdita",
		"zh-cn":  "无损 WAV",
	},
	"profile.default.hq-aac.name": {
		"en-us":  "High-Quality AAC",
		"es-419": "AAC de Alta Calidad",
		"fr-fr":  "AAC de haute qualite",
		"it-it":  "AAC di alta qualita'",
		"zh-cn":  "高质量 AAC",
	},
	"profile.default.alac.name": {
		"en-us":  "Lossless ALAC",
		"es-419": "ALAC sin pérdidas",
		"fr-fr":  "ALAC sans perte",
		"it-it":  "ALAC senza perdita",
		"zh-cn":  "无损 ALAC",
	},

	"shell.tab.syncs": {
		"en-us":  "Syncs",
		"es-419": "Sincronizaciones",
		"fr-fr":  "Synchronisations",
		"it-it":  "Sincronizzazioni",
		"zh-cn":  "同步",
	},
	"shell.tab.profiles": {
		"en-us":  "Profiles",
		"es-419": "Perfiles",
		"fr-fr":  "Profils",
		"it-it":  "Profili",
		"zh-cn":  "配置文件",
	},
	"shell.tab.settings": {
		"en-us":  "Settings",
		"es-419": "Ajustes",
		"fr-fr":  "Options",
		"it-it":  "Opzioni",
		"zh-cn":  "设置",
	},
	"shell.tab.in-progress": {
		"en-us":  "In Progress ($1/$2)",
		"es-419": "En progreso ($1/$2)",
		"fr-fr":  "En cours ($1/$2)",
		"it-it":  "In corso ($1/$2)",
		"zh-cn":  "进行中 ($1/$2)",
	},

	"config.error.unsupported-version": {
		"en-us":  "Unsupported config version",
		"es-419": "Versión de configuración no soportada",
		"fr-fr":  "Version de configuration non-prise en charge",
		"it-it":  "Versione di configurazione non supportata",
		"zh-cn":  "不支持的配置版本",
	},
	"config.error.unknown-profile": {
		"en-us":  "Config contains a reference to an unknown profile",
		"es-419": "Configuración contiene una referencia a un perfil desconocido",
		"fr-fr":  "La configuration contient une reference a un profil inconnu",
		"it-it":  "La configurazione contiene un riferimento a uno profilo sconosciuto",
		"zh-cn":  "配置包含对未知配置文件的引用",
	},

	"setup.title": {
		"en-us":  "Setup",
		"es-419": "Configuración",
		"fr-fr":  "Configuration",
		"it-it":  "Configurazione",
		"zh-cn":  "设置",
	},
	"setup.select-language": {
		"en-us":  "Select Language",
		"es-419": "Seleccionar Idioma",
		"fr-fr":  "Selectionner langue",
		"it-it":  "Selezionare lingua",
		"zh-cn":  "选择语言",
	},
	"setup.setup-complete": {
		"en-us":  "Setup Complete",
		"es-419": "Configuración Completa",
		"fr-fr":  "Configuration terminée",
		"it-it":  "Configurazione Completata",
		"zh-cn":  "设置完成",
	},
	"setup.setup-complete.description": {
		"en-us":  "Setup complete, please restart the application.",
		"es-419": "Configuración completa, reinicie la aplicación.",
		"fr-fr":  "Configuration terminée, veuillez redémarrer l'application.",
		"it-it":  "Configurazione completata, riavviare l'applicazione.",
		"zh-cn":  "设置完成，请重新启动应用程序。",
	},

	"tab.syncs.create": {
		"en-us":  "Create Sync",
		"es-419": "Crear Sincronización",
		"fr-fr":  "Creer synchronisation",
		"it-it":  "Creare sincronizzazione",
		"zh-cn":  "创建同步",
	},
	"tab.syncs.delete-confirm.title": {
		"en-us":  "Delete Sync",
		"es-419": "Eliminar Sincronización",
		"fr-fr":  "Effacer synchonisation",
		"it-it":  "Eliminare sincronizzazione",
		"zh-cn":  "删除同步",
	},
	"tab.syncs.delete-confirm.description": {
		"en-us":  "Are you sure you want to delete the sync \"$1\"?",
		"es-419": "¿Estás seguro de que quieres eliminar la sincronización \"$1\"?",
		"fr-fr":  "Etes-vous certain de vouloir effacer la synchronisation '$1' ?",
		"it-it":  "E' sicura di volere eliminare la sincronizzazione '$1'?",
		"zh-cn":  "您确定要删除同步 \"$1\"？",
	},
	"tab.syncs.form.name": {
		"en-us":  "Name",
		"es-419": "Nombre",
		"fr-fr":  "Nom",
		"it-it":  "Nome",
		"zh-cn":  "名称",
	},
	"tab.syncs.form.source-dir": {
		"en-us":  "Source Directory",
		"es-419": "Directorio de Origen",
		"fr-fr":  "Repertoire source",
		"it-it":  "Cartella sorgente",
		"zh-cn":  "源目录",
	},
	"tab.syncs.form.dest-dir": {
		"en-us":  "Destination Directory",
		"es-419": "Directorio de Destino",
		"fr-fr":  "Repertoire cible",
		"it-it":  "Cartella di destinazione",
		"zh-cn":  "目标目录",
	},
	"tab.syncs.form.profile": {
		"en-us":  "Profile",
		"es-419": "Perfil",
		"fr-fr":  "Profil",
		"it-it":  "Profilo",
		"zh-cn":  "配置文件",
	},
	"tab.syncs.form.skip-files-larger-than": {
		"en-us":  "Skip Files Larger Than",
		"es-419": "Omitir archivos más grandes que",
		"fr-fr":  "Ignorer les fichiers plus gros que",
		"it-it":  "Salta file più grandi di",
		"zh-cn":  "跳过大于",
	},
	"tab.syncs.form.escape-filenames": {
		"en-us":  "Replace invalid characters in filenames?",
		"es-419": "¿Reemplazar caracteres no válidos en los nombres de archivos?",
		"fr-fr":  "Remplacer les caracteres invalides dans les noms de fichiers ?",
		"it-it":  "Sostituire i caratteri invalidi nei nomi dei documenti?",
		"zh-cn":  "替换文件名中的无效字符？",
	},
	"tab.syncs.form.reencode-same-format": {
		"en-us":  "Reencode files with the same format?",
		"es-419": "¿Reencodificar archivos con el mismo formato?",
		"fr-fr":  "Reencoder les fichiers avec le meme format ?",
		"it-it":  "Ricodificare i documenti col stesso formato?",
		"zh-cn":  "重新编码具有相同格式的文件？",
	},
	"tab.syncs.form.error.missing-name": {
		"en-us":  "Name is required",
		"es-419": "Se requiere el nombre",
		"fr-fr":  "Nom manquant",
		"it-it":  "Nome mancante",
		"zh-cn":  "名称为必填项",
	},
	"tab.syncs.form.error.missing-source-dir": {
		"en-us":  "Source directory is required",
		"es-419": "Se requiere el directorio de origen",
		"fr-fr":  "Repertoire source manquant",
		"it-it":  "Cartella sorgente mancante",
		"zh-cn":  "源目录为必填项",
	},
	"tab.syncs.form.error.missing-dest-dir": {
		"en-us":  "Destination directory is required",
		"es-419": "Se requiere el directorio de destino",
		"fr-fr":  "Repertoire cible manquant",
		"it-it":  "Cartella di destinazione mancante",
		"zh-cn":  "目标目录为必填项",
	},
	"tab.syncs.form.error.missing-profile": {
		"en-us":  "Profile is required. Go create a profile first.",
		"es-419": "Se requiere el perfil. Primero cree un perfil.",
		"fr-fr":  "Profil manquant. Creez un profil d'abord.",
		"it-it":  "Profilo mancante. Crei uno profilo prima.",
		"zh-cn":  "配置文件为必填项。首先创建配置文件。",
	},
	"tab.syncs.form.error.name-exists": {
		"en-us":  "A sync with the same name already exists",
		"es-419": "Ya existe una sincronización con el mismo nombre",
		"fr-fr":  "Une synchronisation du meme nom existe deja",
		"it-it":  "Una sincronizzazione col stesso nome esiste gia'",
		"zh-cn":  "同名同步已存在",
	},
	"tab.syncs.form.error.source-dir-nonexistent": {
		"en-us":  "Source directory does not exist, or is not a directory",
		"es-419": "El directorio de origen no existe o no es un directorio",
		"fr-fr":  "Le repertoire source n'existe pas ou n'est pas un repertoire",
		"it-it":  "La cartella sorgente non esiste oppure non e' una cartella",
		"zh-cn":  "源目录不存在或不是目录",
	},
	"tab.syncs.form.error.dest-dir-nonexistent": {
		"en-us":  "Destination directory does not exist, or is not a directory",
		"es-419": "El directorio de destino no existe o no es un directorio",
		"fr-fr":  "Le repertoire cible n'existe pas ou n'est pas un repertoire",
		"it-it":  "La cartella di destinazione non esiste oppure non e' una cartella",
		"zh-cn":  "目标目录不存在或不是目录",
	},
	"tab.syncs.form.error.source-dest-dirs-same": {
		"en-us":  "Source and destination directories cannot be the same",
		"es-419": "Los directorios de origen y destino no pueden ser los mismos",
		"fr-fr":  "Le repertoire source et cible ne peuvent etre le meme repertoire",
		"it-it":  "La cartella sorgente e di destinazione non possono essere la stessa cartella",
		"zh-cn":  "源和目标目录不能相同",
	},
	"tab.syncs.form.error.unable-to-parse-file-size": {
		"en-us":  "Unable to parse file size: $1",
		"es-419": "No se puede analizar el tamaño del archivo: $1",
		"fr-fr":  "Impossible d'analyser la taille du fichier : $1",
		"it-it":  "Impossibile analizzare la dimensione del file: $1",
		"zh-cn":  "无法解析文件大小：$1",
	},
	"tab.syncs.form.warning.destination-fat32-file-size": {
		"en-us":  "The destination directory is on a FAT32 filesystem, which cannot store files larger than 4 GiB.\n\nYou may want to skip files larger than 4 GiB, otherwise there will be errors when syncing.\n\nYour changes have been saved anyway.",
		"es-419": "El directorio de destino está en un sistema de archivos FAT32, que no puede almacenar archivos más grandes de 4 GiB.\n\nEs posible que desee omitir archivos más grandes de 4 GiB, de lo contrario, se producirán errores al sincronizar.\n\nTus cambios se han guardado de todos modos.",
		"fr-fr":  "Le répertoire de destination est sur un système de fichiers FAT32, qui ne peut pas stocker des fichiers plus gros que 4 GiB.\n\nVous pouvez vouloir ignorer les fichiers plus gros que 4 GiB, sinon il y aura des erreurs lors de la synchronisation.\n\nVotre changement a été sauvegardé de toute façon.",
		"it-it":  "La directory di destinazione è su un filesystem FAT32, che non può memorizzare file più grandi di 4 GiB.\n\nPotresti voler saltare i file più grandi di 4 GiB, altrimenti ci saranno errori durante la sincronizzazione.\n\nI tuoi cambiamenti sono stati salvati comunque.",
		"zh-cn":  "目标目录在FAT32文件系统上，它无法存储大于4GiB的文件。\n\n您可能希望跳过大于4GiB的文件，否则在同步时会出现错误。\n\n尽管您的更改已经被保存。",
	},

	"tab.profiles.create": {
		"en-us":  "Create Profile",
		"es-419": "Crear Perfil",
		"fr-fr":  "Creer profil",
		"it-it":  "Creare profilo",
		"zh-cn":  "创建配置文件",
	},
	"tab.profiles.form.name": {
		"en-us":  "Name",
		"es-419": "Nombre",
		"fr-fr":  "Nom",
		"it-it":  "Nome",
		"zh-cn":  "名称",
	},
	"tab.profiles.form.format": {
		"en-us":  "Format",
		"es-419": "Formato",
		"fr-fr":  "Format",
		"it-it":  "Formato",
		"zh-cn":  "格式",
	},
	"tab.profiles.form.is-lossless": {
		"en-us":  "Lossless",
		"es-419": "Sin pérdidas",
		"fr-fr":  "Sans perte",
		"it-it":  "Senza perdita",
		"zh-cn":  "无损",
	},
	"tab.profiles.form.supports-metadata": {
		"en-us":  "Supports Metadata",
		"es-419": "Soporta metadatos",
		"fr-fr":  "Prise en charge pour metadonnees",
		"it-it":  "Supporta metadati",
		"zh-cn":  "支持元数据",
	},
	"tab.profiles.form.supports-artwork": {
		"en-us":  "Supports Artwork",
		"es-419": "Soporta arte",
		"fr-fr":  "Prise en charge pour illustrations",
		"it-it":  "Supporta illustrazioni",
		"zh-cn":  "支持艺术",
	},
	"tab.profiles.form.bitrate": {
		"en-us":  "Bitrate",
		"es-419": "Bitrate",
		"fr-fr":  "Debit binaire",
		"it-it":  "Velocita' di trasmissione",
		"zh-cn":  "比特率",
	},
	"tab.profiles.form.error.invalid-bitrate": {
		"en-us":  "Invalid bitrate",
		"es-419": "Bitrate inválido",
		"fr-fr":  "Debit binaire invalide",
		"it-it":  "Velocita' di trasmissione invalida",
		"zh-cn":  "无效的比特率",
	},
	"tab.profiles.form.error.missing-name": {
		"en-us":  "Name is required",
		"es-419": "Se requiere el nombre",
		"fr-fr":  "Nom manquant",
		"it-it":  "Nome mancante",
		"zh-cn":  "名称为必填项",
	},
	"tab.profiles.form.error.name-exists": {
		"en-us":  "A profile with the same name already exists",
		"es-419": "Ya existe un perfil con el mismo nombre",
		"fr-fr":  "Un profil du meme nom existe deja",
		"it-it":  "Un profilo col stesso nome esiste gia'",
		"zh-cn":  "同名配置文件已存在",
	},
	"tab.profiles.delete-confirm.title": {
		"en-us":  "Delete Profile",
		"es-419": "Eliminar Perfil",
		"fr-fr":  "Effacer profil",
		"it-it":  "Eliminare profilo",
		"zh-cn":  "删除配置文件",
	},
	"tab.profiles.delete-confirm.description": {
		"en-us":  "Are you sure you want to delete the profile \"$1\"?",
		"es-419": "¿Estás seguro de que quieres eliminar el perfil \"$1\"?",
		"fr-fr":  "Etes-vous certain de vouloir effacer le profil '$1' ?",
		"it-it":  "E' sicura di volere eliminare il profilo '$1'?",
		"zh-cn":  "您确定要删除配置文件 \"$1\"？",
	},
	"tab.profiles.error.in-use": {
		"en-us":  "Profile is in use by a sync. Delete syncs using this profile first.",
		"es-419": "El perfil está en uso por una sincronización. Elimina primero las sincronizaciones que usan este perfil.",
		"fr-fr":  "Ce profil est utilise par une synchronisation. Effacez d'abord les synchronisations utiliaant ce profil.",
		"it-it":  "Questo profilo e' gia' usato da una sincronizzazione. Elimini prima le sincronizzazioni usando questo profilo.",
		"zh-cn":  "配置文件正在被同步使用。先删除同步使用的配置文件。",
	},

	"tab.progress.start": {
		"en-us":  "Start",
		"es-419": "Iniciar",
		"fr-fr":  "Commencer",
		"it-it":  "Cominciare",
		"zh-cn":  "开始",
	},
	"tab.progress.cancel": {
		"en-us":  "Cancel",
		"es-419": "Cancelar",
		"fr-fr":  "Annuler",
		"it-it":  "Annullare",
		"zh-cn":  "取消",
	},
	"tab.progress.status-label": {
		"en-us":  "Completed: $1/$2, skipped: $3, failed: $4",
		"es-419": "Completado: $1/$2, omitidos: $3, fallidos: $4",
		"fr-fr":  "Succes : $1/$2, omission(s) : $3, echec(s) : $4",
		"it-it":  "Riuscito/i: $1/$2, omesso/i: $3, fallimento/i: $4",
		"zh-cn":  "已完成: $1/$2, 跳过: $3, 失败: $4",
	},

	"sync.scanning-source": {
		"en-us":  "Scanning source directory...",
		"es-419": "Escaneando directorio de origen...",
		"fr-fr":  "Examination du repertoire source en cours...",
		"it-it":  "Esaminando della cartella sorgente...",
		"zh-cn":  "正在扫描源目录...",
	},
	"sync.copying": {
		"en-us":  "Copying $1",
		"es-419": "Copiando $1",
		"fr-fr":  "Copie de $1",
		"it-it":  "Copiando $1",
		"zh-cn":  "正在复制 $1",
	},
	"sync.transcoding": {
		"en-us":  "Transcoding $1",
		"es-419": "Transcodificando $1",
		"fr-fr":  "Transcodage de $1",
		"it-it":  "Transcodificando $1",
		"zh-cn":  "正在转码 $1",
	},
	"sync.destination-not-found": {
		"en-us":  "Destination directory not found, is it on a disconnected drive?",
		"es-419": "Directorio de destino no encontrado, ¿está en una unidad desconectada?",
		"fr-fr":  "Le répertoire de destination n'a pas été trouvé, est-il sur un lecteur déconnecté ?",
		"it-it":  "Cartella di destinazione non trovata, è presente su un disco disconnesso?",
		"zh-cn":  "目标目录未找到，是否在断开的驱动器上？",
	},
	"sync.error-checking-destination-directory": {
		"en-us":  "Error checking destination directory, is your computer allowed to access it? Error: $1",
		"es-419": "Error al comprobar el directorio de destino, ¿está permitido el acceso a él?",
		"fr-fr":  "Erreur lors de la vérification du répertoire de destination, est-ce que votre ordinateur est autorisé à y accéder ?",
		"it-it":  "Errore durante il controllo della cartella di destinazione, è consentito l'accesso a essa?",
		"zh-cn":  "检查目标目录时出错，您的计算机是否允许访问它？",
	},
	"sync.path-already-exists": {
		"en-us":  "Path $1 already exists, skipping",
		"es-419": "La ruta $1 ya existe, se omite",
		"fr-fr":  "Le chemin d'acces $1 existe deja, omission de celui-ci",
		"it-it":  "Il percorso $1 esiste gia', omissione di quello",
		"zh-cn":  "路径 $1 已存在，将跳过",
	},
	"sync.error-processing-file-x-y": {
		"en-us":  "Error processing path \"$1\": $2",
		"es-419": "Error procesando la ruta \"$1\": $2",
		"fr-fr":  "Erreur lors du traitement du chemin \"$1\" : $2",
		"it-it":  "Errore durante l'elaborazione del percorso \"$1\": $2",
		"zh-cn":  "处理路径“$1”时出错：$2",
	},
	"sync.could-not-find-executable-path": {
		"en-us":  "Could not find executable path, therefore cannot locate bundled FFmpeg binaries. Error: $1",
		"es-419": "No se pudo encontrar la ruta del ejecutable, por lo tanto, no se pueden localizar los binarios FFmpeg empaquetados. Error: $1",
		"fr-fr":  "Impossible de trouver le chemin d'accès de l'exécutable, par conséquent, les binaires FFmpeg empaquetés ne peuvent pas être localisés. Erreur : $1",
		"it-it":  "Impossibile trovare il percorso dell'eseguibile, pertanto, non è possibile localizzare i binari FFmpeg bundlati. Errore: $1",
		"zh-cn":  "无法找到可执行文件路径，因此无法定位包装的 FFmpeg 二进制文件。错误：$1",
	},
	"sync.could-not-locate-bundled-ffmpeg-binaries": {
		"en-us":  "Could not locate bundled FFmpeg binaries, falling back to system-wide FFmpeg installation. Conversion may fail!",
		"es-419": "No se pudieron localizar los binarios FFmpeg empaquetados, regresando a la instalación global de FFmpeg del sistema. ¡La conversión puede fallar!",
		"fr-fr":  "Impossible de localiser les binaires FFmpeg empaquetés, retour à l'installation globale FFmpeg du système. La conversion peut échouer !",
		"it-it":  "Impossibile localizzare i binari FFmpeg bundlati, fallendo sull'installazione globale di FFmpeg del sistema. La conversione potrebbe fallire!",
		"zh-cn":  "无法定位包装的 FFmpeg 二进制文件，正在回退到系统全局的 FFmpeg 安装。转换可能会失败！",
	},
	"sync.done": {
		"en-us":  "Done (total: $1, completed: $2, skipped: $3, failed: $4)",
		"es-419": "Hecho (total: $1, completados: $2, omitidos: $3, fallidos: $4)",
		"fr-fr":  "Fini (total : $1, succes : $2, omission(s) : $3, echec(s) : $4)",
		"it-it":  "Finito (totale: $1, riuscito/i: $2, omesso/i: $3, fallimento/i: $4)",
		"zh-cn":  "完成 (总计: $1, 已完成: $2, 跳过: $3, 失败: $4)",
	},

	"parse.error.number-cannot-be-negative": {
		"en-us":  "Number cannot be negative",
		"es-419": "El número no puede ser negativo",
		"fr-fr":  "Le nombre ne peut pas être négatif",
		"it-it":  "Il numero non può essere negativo",
		"zh-cn":  "数字不能为负数",
	},
	"parse.error.missing-unit-suffix": {
		"en-us":  "Missing unit suffix (such as MiB or GiB)",
		"es-419": "Falta el sufijo de unidad (como MiB o GiB)",
		"fr-fr":  "Suffixe d'unité manquant (tel que MiB ou GiB)",
		"it-it":  "Manca il suffisso di unità (come MiB o GiB)",
		"zh-cn":  "缺少单位后缀（例如MiB或GiB）",
	},
	"parse.error.missing-number": {
		"en-us":  "Missing number",
		"es-419": "Falta el número",
		"fr-fr":  "Nombre manquant",
		"it-it":  "Manca il numero",
		"zh-cn":  "缺少数字",
	},
	"parse.error.please-use-dot-for-decimal-points": {
		"en-us":  "Please use a \".\" for decimal points",
		"es-419": "Por favor, use un \".\" para los puntos decimales",
		"fr-fr":  "Veuillez utiliser un \".\" pour les points décimaux",
		"it-it":  "Utilizza un \".\" per i punti decimali",
		"zh-cn":  "请使用“.”作为小数点",
	},
	"parse.error.unknown-unit-x": {
		"en-us":  "Unknown unit \"$1\"",
		"es-419": "Unidad desconocida \"$1\"",
		"fr-fr":  "Unité inconnue \"$1\"",
		"it-it":  "Unità sconosciuta \"$1\"",
		"zh-cn":  "未知单位“$1”",
	},
}
