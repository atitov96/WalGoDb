# go.mod

```mod
module atitov96/walgodb

go 1.23.2

require (
	github.com/stretchr/testify v1.8.1
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

```

# config.yaml

```yaml
engine:
  type: "in_memory"
network:
  address: "127.0.0.1:3223"
  max_connections: 100
  max_message_size: "4KB"
  idle_timeout: 5s
logging:
  level: "info"
#  output: "/log/output.log"
```

# README.md

```md
# WalGoDb

### Синтаксис запросов

Грамматика языка запросов в виде eBNF:
\`\`\`
query = set_command | get_command | del_command

set_command = "SET" argument argument
get_command = "GET" argument
del_command = "DEL" argument
argument    = punctuation | letter | digit { punctuation | letter | digit }

punctuation = "*" | "/" | "_" | ...
letter      = "a" | ... | "z" | "A" | ... | "Z"
digit       = "0" | ... | "9"
\`\`\`

### Пример запросов:
\`\`\`
SET weather_2_pm cold_moscow_weather
GET /etc/nginx/config
DEL user_****
\`\`\`

То есть, синтаксис очень простой - есть всего лишь три возможных команды (SET, GET, DEL - регистрозависимые), аргументами команд являются только следующая возможная комбинация /(\w+)/g , разделителями являются любые пробельные символы.



### Особенности реализации

Взаимодействие с базой данных на данном этапе должно осуществляться только лишь через командную строку, то есть вы запускаете базу данных и взаимодействуете с ней через командую строку с использованием языка запросов. Никакие данные на жесткий диск сейчас дампить не нужно (если вы сделали рестарт базы данных, то все данные должны будут пропасть)

В качестве engine пишем in-memory движок, который просто будет хранить данные в виде хэш-таблицы (в качестве ключей и запросов только текстовые типы данных).

Необходимо сразу задуматься над логированием, поэтому следует логировать ключевые моменты обработки запросов (в качестве библиотеки логирования можно выбрать zap)

Покрываем код тестами, особенно это касается компонентов parser, analyzer внутри compute слоя и компонента engine внутри storage слоя (суммарное покрытие кода тестами должно быть не менее 90%)

### Домашнее задание №2

Необходимо реализовать TCP сервер для базы данных (внутри существующего приложения), чтобы с ним можно было бы взаимодействовать по сети.

Протокол общения с сервером базы данных будет текстовым, то есть запросы будут передаваться в точно таком же формате, как и до этого вы вводили в командной строке.

Также необходимо реализовать отдельное приложение CLI клиент (можно в виде отдельного main.go файла), с помощью которого вы будете дальше взаимодействовать с вашей базой данных. По сути это приложение должно состоять только лишь из TCP клиента и удобного интерфейса командной строки для работы с вашей базой данных.

### Конфигурация

В этом домашнем задании вам нужно реализовать возможность конфигурации вашей базы данных из yaml файла. В этом файле можно будет указать параметры сетевого соединения, вид движка (у нас только in_memory, но тем не менее), а также уровень логирования. В будущем эта конфигурация будет только расширяться. Пример файла конфигурации:

\`\`\`yaml
engine:
  type: "in_memory"
network:
  address: "127.0.0.1:3223"
  max_connections: 100
  max_message_size: "4KB"
  idle_timeout: 5m
logging:
  level: "info"
  output: "/log/output.log"
\`\`\`

### Особенности реализации

* Каждый клиент должен обрабатываться в отдельной горутине, важно не забыть про синхронизацию работы этих клиентов с хэш-таблицей, где хранятся данные
* Следует реализовать ограничитель количества одновременных соединений с вашим сервером базы данных - это значит, что ваш сервер должен уметь одновременно обрабатывать не больше, чем N клиентов
* Следует реализовать конфигурирование вашей базы данных таким образом, что если не задан какой-либо параметр конфигурации, то должны использоваться параметры по умолчанию, а не завершать приложение с ошибкой конфигурации
* Следует реализовать возможность конфигурирования CLI клиента с использованием аргументов командной строки, например вот так: ./database_client —-address=localhost:3223
* Покрываем код тестами, особенно это касается слоя сетевого взаимодействия и конфигурирования базы данных (суммарное покрытие кода тестами должно быть не менее 90%)

```

# .gitignore

```
# Ignore .idea folder (JetBrains IDEs)
.idea/

# Binaries for programs and plugins
*.exe
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Directories created by the "go get" command
/vendor/

# Dependency directories (remove the comment below if you use Go modules)
# /vendor/

# Go workspace file
go.work
go.work.sum

# Exclude Go build cache
/go/pkg/

# IDE-specific ignores
# Visual Studio Code
.vscode/
.history/
```

# .idea/workspace.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="AutoImportSettings">
    <option name="autoReloadType" value="ALL" />
  </component>
  <component name="ChangeListManager">
    <list default="true" id="36068d6a-4ec3-4d78-9c6f-6e43182855b5" name="Changes" comment="">
      <change afterPath="$PROJECT_DIR$/cmd/client/main.go" afterDir="false" />
      <change beforePath="$PROJECT_DIR$/README.md" beforeDir="false" afterPath="$PROJECT_DIR$/README.md" afterDir="false" />
      <change beforePath="$PROJECT_DIR$/cmd/main.go" beforeDir="false" />
      <change beforePath="$PROJECT_DIR$/go.mod" beforeDir="false" afterPath="$PROJECT_DIR$/go.mod" afterDir="false" />
      <change beforePath="$PROJECT_DIR$/internal/compute/compute.go" beforeDir="false" afterPath="$PROJECT_DIR$/internal/compute/compute.go" afterDir="false" />
      <change beforePath="$PROJECT_DIR$/internal/compute/compute_test.go" beforeDir="false" afterPath="$PROJECT_DIR$/internal/compute/compute_test.go" afterDir="false" />
      <change beforePath="$PROJECT_DIR$/pkg/logger/logger.go" beforeDir="false" afterPath="$PROJECT_DIR$/pkg/logger/logger.go" afterDir="false" />
    </list>
    <option name="SHOW_DIALOG" value="false" />
    <option name="HIGHLIGHT_CONFLICTS" value="true" />
    <option name="HIGHLIGHT_NON_ACTIVE_CHANGELIST" value="false" />
    <option name="LAST_RESOLUTION" value="IGNORE" />
  </component>
  <component name="FileTemplateManagerImpl">
    <option name="RECENT_TEMPLATES">
      <list>
        <option value="Go File" />
      </list>
    </option>
  </component>
  <component name="GOROOT" url="file:///opt/homebrew/opt/go/libexec" />
  <component name="Git.Settings">
    <option name="RECENT_GIT_ROOT_PATH" value="$PROJECT_DIR$" />
  </component>
  <component name="ProjectColorInfo">{
  &quot;associatedIndex&quot;: 5
}</component>
  <component name="ProjectId" id="2lwVo27H45j7sJyjJRGjnExhq7j" />
  <component name="ProjectViewState">
    <option name="hideEmptyMiddlePackages" value="true" />
    <option name="showLibraryContents" value="true" />
  </component>
  <component name="PropertiesComponent"><![CDATA[{
  "keyToString": {
    "DefaultGoTemplateProperty": "Go File",
    "Go Build.go build atitov96/walgodb/cmd.executor": "Run",
    "Go Build.go build atitov96/walgodb/cmd/client.executor": "Run",
    "Go Build.go build atitov96/walgodb/cmd/server.executor": "Run",
    "Go Test.BenchmarkDel in atitov96/walgodb/internal/compute.executor": "Profiler#3",
    "Go Test.BenchmarkGet in atitov96/walgodb/internal/compute.executor": "Run",
    "Go Test.BenchmarkSet in atitov96/walgodb/internal/compute.executor": "Run",
    "Go Test.TestComputeLayer_Execute in atitov96/walgodb/internal/compute.executor": "Run",
    "Go Test.TestComputeLayer_Execute/GET_non-existing_key in atitov96/walgodb/internal/compute.executor": "Run",
    "Go Test.TestComputeLayer_Execute/Invalid_command in atitov96/walgodb/internal/compute.executor": "Coverage",
    "Go Test.TestComputeLayer_Execute/SET_key in atitov96/walgodb/internal/compute.executor": "Debug",
    "Go Test.TestInMemoryEngine in atitov96/walgodb/internal/compute.executor": "Run",
    "Go Test.TestLoadConfig in atitov96/walgodb/pkg/config.executor": "Coverage",
    "Go Test.TestLoadConfigDefaults in atitov96/walgodb/pkg/config.executor": "Coverage",
    "Go Test.TestParser_Parse in atitov96/walgodb/internal/compute.executor": "Run",
    "Go Test.TestParser_Parse/Invalid_arguments in atitov96/walgodb/internal/compute.executor": "Run",
    "Go Test.TestServer in atitov96/walgodb/cmd/server.executor": "Coverage",
    "Go Test.go test atitov96/walgodb/internal/compute.executor": "Coverage",
    "Go Test.go test atitov96/walgodb/pkg/config.executor": "Coverage",
    "Go Test.gobench atitov96/walgodb/internal/compute.executor": "Run",
    "RunOnceActivity.ShowReadmeOnStart": "true",
    "RunOnceActivity.go.formatter.settings.were.checked": "true",
    "RunOnceActivity.go.migrated.go.modules.settings": "true",
    "RunOnceActivity.go.modules.automatic.dependencies.download": "true",
    "RunOnceActivity.go.modules.go.list.on.any.changes.was.set": "true",
    "git-widget-placeholder": "iteration-1",
    "go.import.settings.migrated": "true",
    "go.sdk.automatically.set": "true",
    "last_opened_file_path": "/Users/titoffag/Documents/WalGoDb",
    "node.js.detected.package.eslint": "true",
    "node.js.detected.package.tslint": "true",
    "node.js.selected.package.eslint": "(autodetect)",
    "node.js.selected.package.tslint": "(autodetect)",
    "nodejs_package_manager_path": "npm",
    "settings.editor.selected.configurable": "preferences.pluginManager"
  }
}]]></component>
  <component name="RecentsManager">
    <key name="MoveFile.RECENT_KEYS">
      <recent name="$PROJECT_DIR$/cmd/server" />
      <recent name="$PROJECT_DIR$/cmd/client" />
      <recent name="$PROJECT_DIR$/pkg" />
    </key>
  </component>
  <component name="RunManager" selected="Go Test.TestServer in atitov96/walgodb/cmd/server">
    <configuration name="go build atitov96/walgodb/cmd/client" type="GoApplicationRunConfiguration" factoryName="Go Application" temporary="true" nameIsGenerated="true">
      <module name="WalGoDb" />
      <working_directory value="$PROJECT_DIR$" />
      <kind value="PACKAGE" />
      <package value="atitov96/walgodb/cmd/client" />
      <directory value="$PROJECT_DIR$" />
      <filePath value="$PROJECT_DIR$/cmd/client/main.go" />
      <method v="2" />
    </configuration>
    <configuration name="TestLoadConfig in atitov96/walgodb/pkg/config" type="GoTestRunConfiguration" factoryName="Go Test" temporary="true" nameIsGenerated="true">
      <module name="WalGoDb" />
      <working_directory value="$PROJECT_DIR$/pkg/config" />
      <root_directory value="$PROJECT_DIR$" />
      <kind value="PACKAGE" />
      <package value="atitov96/walgodb/pkg/config" />
      <directory value="$PROJECT_DIR$" />
      <filePath value="$PROJECT_DIR$" />
      <framework value="gotest" />
      <pattern value="^\QTestLoadConfig\E$" />
      <method v="2" />
    </configuration>
    <configuration name="TestLoadConfigDefaults in atitov96/walgodb/pkg/config" type="GoTestRunConfiguration" factoryName="Go Test" temporary="true" nameIsGenerated="true">
      <module name="WalGoDb" />
      <working_directory value="$PROJECT_DIR$/pkg/config" />
      <root_directory value="$PROJECT_DIR$" />
      <kind value="PACKAGE" />
      <package value="atitov96/walgodb/pkg/config" />
      <directory value="$PROJECT_DIR$" />
      <filePath value="$PROJECT_DIR$" />
      <framework value="gotest" />
      <pattern value="^\QTestLoadConfigDefaults\E$" />
      <method v="2" />
    </configuration>
    <configuration name="TestServer in atitov96/walgodb/cmd/server" type="GoTestRunConfiguration" factoryName="Go Test" temporary="true" nameIsGenerated="true">
      <module name="WalGoDb" />
      <working_directory value="$PROJECT_DIR$/cmd/server" />
      <root_directory value="$PROJECT_DIR$" />
      <kind value="PACKAGE" />
      <package value="atitov96/walgodb/cmd/server" />
      <directory value="$PROJECT_DIR$" />
      <filePath value="$PROJECT_DIR$" />
      <framework value="gotest" />
      <pattern value="^\QTestServer\E$" />
      <method v="2" />
    </configuration>
    <configuration name="go test atitov96/walgodb/pkg/config" type="GoTestRunConfiguration" factoryName="Go Test" temporary="true" nameIsGenerated="true">
      <module name="WalGoDb" />
      <working_directory value="$PROJECT_DIR$/pkg/config" />
      <root_directory value="$PROJECT_DIR$" />
      <kind value="PACKAGE" />
      <package value="atitov96/walgodb/pkg/config" />
      <directory value="$PROJECT_DIR$" />
      <filePath value="$PROJECT_DIR$" />
      <framework value="gotest" />
      <method v="2" />
    </configuration>
    <recent_temporary>
      <list>
        <item itemvalue="Go Test.TestServer in atitov96/walgodb/cmd/server" />
        <item itemvalue="Go Test.go test atitov96/walgodb/pkg/config" />
        <item itemvalue="Go Test.TestLoadConfig in atitov96/walgodb/pkg/config" />
        <item itemvalue="Go Test.TestLoadConfigDefaults in atitov96/walgodb/pkg/config" />
        <item itemvalue="Go Build.go build atitov96/walgodb/cmd/client" />
      </list>
    </recent_temporary>
  </component>
  <component name="SharedIndexes">
    <attachedChunks>
      <set>
        <option value="bundled-gosdk-5df93f7ad4aa-df9ad98b711f-org.jetbrains.plugins.go.sharedIndexes.bundled-GO-242.22855.106" />
        <option value="bundled-js-predefined-d6986cc7102b-5c90d61e3bab-JavaScript-GO-242.22855.106" />
      </set>
    </attachedChunks>
  </component>
  <component name="SpellCheckerSettings" RuntimeDictionaries="0" Folders="0" CustomDictionaries="0" DefaultDictionary="application-level" UseSingleDictionary="true" transferred="true" />
  <component name="TypeScriptGeneratedFilesManager">
    <option name="version" value="3" />
  </component>
  <component name="VgoProject">
    <settings-migrated>true</settings-migrated>
  </component>
  <component name="com.intellij.coverage.CoverageDataManagerImpl">
    <SUITE FILE_PATH="coverage/WalGoDb$go_test_atitov96_walgodb_pkg_config.out" NAME="go test atitov96/walgodb/pkg/config Coverage Results" MODIFIED="1728054784739" SOURCE_PROVIDER="com.intellij.coverage.DefaultCoverageFileProvider" RUNNER="GoCoverage" COVERAGE_BY_TEST_ENABLED="false" COVERAGE_TRACING_ENABLED="false" />
    <SUITE FILE_PATH="coverage/WalGoDb$TestLoadConfig_in_atitov96_walgodb_pkg_config.out" NAME="TestLoadConfig in atitov96/walgodb/pkg/config Coverage Results" MODIFIED="1728054768648" SOURCE_PROVIDER="com.intellij.coverage.DefaultCoverageFileProvider" RUNNER="GoCoverage" COVERAGE_BY_TEST_ENABLED="false" COVERAGE_TRACING_ENABLED="false" />
    <SUITE FILE_PATH="coverage/WalGoDb$BenchmarkDel_in_atitov96_walgodb_internal_compute.out" NAME="BenchmarkDel in atitov96/walgodb/internal/compute Coverage Results" MODIFIED="1727984174209" SOURCE_PROVIDER="com.intellij.coverage.DefaultCoverageFileProvider" RUNNER="GoCoverage" COVERAGE_BY_TEST_ENABLED="false" COVERAGE_TRACING_ENABLED="false" />
    <SUITE FILE_PATH="coverage/WalGoDb$TestServer_in_atitov96_walgodb_cmd_server.out" NAME="TestServer in atitov96/walgodb/cmd/server Coverage Results" MODIFIED="1728055400479" SOURCE_PROVIDER="com.intellij.coverage.DefaultCoverageFileProvider" RUNNER="GoCoverage" COVERAGE_BY_TEST_ENABLED="false" COVERAGE_TRACING_ENABLED="false" />
    <SUITE FILE_PATH="coverage/WalGoDb$go_test_atitov96_walgodb_internal_compute.out" NAME="go test atitov96/walgodb/internal/compute Coverage Results" MODIFIED="1727984713395" SOURCE_PROVIDER="com.intellij.coverage.DefaultCoverageFileProvider" RUNNER="GoCoverage" COVERAGE_BY_TEST_ENABLED="false" COVERAGE_TRACING_ENABLED="false" />
    <SUITE FILE_PATH="coverage/WalGoDb$TestComputeLayer_Execute_Invalid_command_in_atitov96_walgodb_internal_compute.out" NAME="TestComputeLayer_Execute/Invalid_command in atitov96/walgodb/internal/compute Coverage Results" MODIFIED="1727984655171" SOURCE_PROVIDER="com.intellij.coverage.DefaultCoverageFileProvider" RUNNER="GoCoverage" COVERAGE_BY_TEST_ENABLED="false" COVERAGE_TRACING_ENABLED="false" />
    <SUITE FILE_PATH="coverage/WalGoDb$TestLoadConfigDefaults_in_atitov96_walgodb_pkg_config.out" NAME="TestLoadConfigDefaults in atitov96/walgodb/pkg/config Coverage Results" MODIFIED="1728054741290" SOURCE_PROVIDER="com.intellij.coverage.DefaultCoverageFileProvider" RUNNER="GoCoverage" COVERAGE_BY_TEST_ENABLED="false" COVERAGE_TRACING_ENABLED="false" />
  </component>
</project>
```

# .idea/vcs.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="VcsDirectoryMappings">
    <mapping directory="" vcs="Git" />
  </component>
</project>
```

# .idea/modules.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="ProjectModuleManager">
    <modules>
      <module fileurl="file://$PROJECT_DIR$/.idea/WalGoDb.iml" filepath="$PROJECT_DIR$/.idea/WalGoDb.iml" />
    </modules>
  </component>
</project>
```

# .idea/material_theme_project_new.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="MaterialThemeProjectNewConfig">
    <option name="metadata">
      <MTProjectMetadataState>
        <option name="migrated" value="true" />
        <option name="pristineConfig" value="false" />
        <option name="userId" value="261fcf15:19198c78901:-7ffe" />
      </MTProjectMetadataState>
    </option>
  </component>
</project>
```

# .idea/WalGoDb.iml

```iml
<?xml version="1.0" encoding="UTF-8"?>
<module type="WEB_MODULE" version="4">
  <component name="Go" enabled="true" />
  <component name="NewModuleRootManager">
    <content url="file://$MODULE_DIR$" />
    <orderEntry type="inheritedJdk" />
    <orderEntry type="sourceFolder" forTests="false" />
  </component>
</module>
```

# .idea/.gitignore

```
# Default ignored files
/shelf/
/workspace.xml
# Editor-based HTTP Client requests
/httpRequests/
# Datasource local storage ignored files
/dataSources/
/dataSources.local.xml

```

# pkg/logger/logger.go

```go
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string, output string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	logLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	config.Level = zap.NewAtomicLevelAt(logLevel)

	if output != "stdout" {
		config.OutputPaths = []string{output}
	}

	logger, err := config.Build()
	return logger, err
}

```

# pkg/config/config_test.go

```go
package config

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLoadConfig(t *testing.T) {
	configContent := `
engine:
  type: "in_memory"
network:
  address: "127.0.0.1:3223"
  max_connections: 100
  max_message_size: "4KB"
  idle_timeout: 5m
logging:
  level: "info"
  output: "stdout"
`
	tmpfile, err := ioutil.TempFile("", "config.*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(configContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	config, err := LoadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.Engine.Type != "in_memory" {
		t.Errorf("Expected engine type 'in_memory', got '%s'", config.Engine.Type)
	}

	if config.Network.Address != "127.0.0.1:3223" {
		t.Errorf("Expected address '127.0.0.1:3223', got '%s'", config.Network.Address)
	}

	if config.Network.MaxConnections != 100 {
		t.Errorf("Expected max connections 100, got %d", config.Network.MaxConnections)
	}

	if config.Network.MaxMessageSize != "4KB" {
		t.Errorf("Expected max message size '4KB', got '%s'", config.Network.MaxMessageSize)
	}

	if config.Network.IdleTimeout != 5*time.Minute {
		t.Errorf("Expected idle timeout 5m, got %v", config.Network.IdleTimeout)
	}

	if config.Logging.Level != "info" {
		t.Errorf("Expected logging level 'info', got '%s'", config.Logging.Level)
	}

	if config.Logging.Output != "stdout" {
		t.Errorf("Expected logging output 'stdout', got '%s'", config.Logging.Output)
	}
}

func TestLoadConfigDefaults(t *testing.T) {
	configContent := `{}` // Empty config

	tmpfile, err := ioutil.TempFile("", "config.*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(configContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	config, err := LoadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.Engine.Type != "in_memory" {
		t.Errorf("Expected default engine type 'in_memory', got '%s'", config.Engine.Type)
	}

	if config.Network.Address != "127.0.0.1:3223" {
		t.Errorf("Expected default address '127.0.0.1:3223', got '%s'", config.Network.Address)
	}

	if config.Network.MaxConnections != 10 {
		t.Errorf("Expected default max connections 100, got %d", config.Network.MaxConnections)
	}

	if config.Network.MaxMessageSize != "4KB" {
		t.Errorf("Expected default max message size '4KB', got '%s'", config.Network.MaxMessageSize)
	}

	if config.Network.IdleTimeout != 5*time.Second {
		t.Errorf("Expected default idle timeout 5m, got %v", config.Network.IdleTimeout)
	}

	if config.Logging.Level != "info" {
		t.Errorf("Expected default logging level 'info', got '%s'", config.Logging.Level)
	}

	if config.Logging.Output != "stdout" {
		t.Errorf("Expected default logging output 'stdout', got '%s'", config.Logging.Output)
	}
}

```

# pkg/config/config.go

```go
package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Engine  EngineConfig  `yaml:"engine"`
	Network NetworkConfig `yaml:"network"`
	Logging LoggingConfig `yaml:"logging"`
}

type EngineConfig struct {
	Type string `yaml:"type"`
}

type NetworkConfig struct {
	Address        string        `yaml:"address"`
	MaxConnections int           `yaml:"max_connections"`
	MaxMessageSize string        `yaml:"max_message_size"`
	IdleTimeout    time.Duration `yaml:"idle_timeout"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	setDefault(&cfg)

	return &cfg, nil
}

func setDefault(config *Config) {
	if config.Engine.Type == "" {
		config.Engine.Type = "in_memory"
	}
	if config.Network.Address == "" {
		config.Network.Address = "127.0.0.1:3223"
	}
	if config.Network.MaxConnections == 0 {
		config.Network.MaxConnections = 10
	}
	if config.Network.MaxMessageSize == "" {
		config.Network.MaxMessageSize = "4KB"
	}
	if config.Network.IdleTimeout == 0 {
		config.Network.IdleTimeout = 5 * time.Second
	}
	if config.Logging.Level == "" {
		config.Logging.Level = "info"
	}
	if config.Logging.Output == "" {
		config.Logging.Output = "stdout"
	}
}

```

# internal/storage/engine.go

```go
package storage

import "sync"

type Storage interface {
	Set(key string, value string)
	Get(key string) (string, bool)
	Delete(key string)
}

type inMemoryEngine struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewInMemoryEngine() Storage {
	return &inMemoryEngine{
		data: make(map[string]string),
	}
}

func (e *inMemoryEngine) Set(key string, value string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.data[key] = value
}

func (e *inMemoryEngine) Get(key string) (string, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	val, ok := e.data[key]
	return val, ok
}

func (e *inMemoryEngine) Delete(key string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.data, key)
}

```

# internal/compute/parser.go

```go
package compute

import (
	"errors"
	"regexp"
	"strings"
)

type parser struct{}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) Parse(expression string) (Command, error) {
	parts := strings.Fields(expression)
	if len(parts) == 0 {
		return Command{}, errors.New("empty expression")
	}
	commandType := strings.ToUpper(parts[0])
	args := parts[1:]

	if !isValidCommandType(commandType) {
		return Command{}, errors.New("unknown command")
	}

	if !areValidArguments(args) {
		return Command{}, errors.New("invalid arguments: must be alphanumeric")
	}

	return Command{Type: commandType, Args: args}, nil
}

func isValidCommandType(commandType string) bool {
	return commandType == "SET" || commandType == "GET" || commandType == "DEL"
}

func areValidArguments(args []string) bool {
	for _, arg := range args {
		if !isAlphanumeric(arg) {
			return false
		}
	}
	return true
}

func isAlphanumeric(s string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9_\\-]*$").MatchString(s)
}

```

# internal/compute/compute_test.go

```go
package compute

import (
	"atitov96/walgodb/internal/storage"
	"atitov96/walgodb/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	p := NewParser()

	tests := []struct {
		name       string
		expression string
		want       Command
		wantErr    bool
	}{
		{"Valid SET command", "SET key value", Command{Type: "SET", Args: []string{"key", "value"}}, false},
		{"Valid GET command", "GET key", Command{Type: "GET", Args: []string{"key"}}, false},
		{"Valid DEL command", "DEL key", Command{Type: "DEL", Args: []string{"key"}}, false},
		{"Empty expression", "", Command{}, true},
		{"Unknown command", "UNKNOWN key", Command{}, true},
		{"Invalid arguments", "SET key! value!", Command{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(tt.expression)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestComputeLayer_Execute(t *testing.T) {
	log, _ := logger.NewLogger("info", "stdout")
	parser := NewParser()
	engine := storage.NewInMemoryEngine()
	compute := NewComputeLayer(parser, engine, log)

	tests := []struct {
		name    string
		command string
		want    string
		wantErr bool
	}{
		{"SET key", "SET test_key test_value", "OK", false},
		{"GET existing key", "GET test_key", "test_value", false},
		{"GET non-existing key", "GET non_existing_key", "NOT FOUND", false},
		{"DEL existing key", "DEL test_key", "OK", false},
		{"DEL non-existing key", "DEL non_existing_key", "OK", false},
		{"Invalid command", "INVALID test_key", "", true},
		{"Set with wrong number of arguments", "SET test_key", "", true},
		{"Get with wrong number of arguments", "GET key1 key2", "", true},
		{"Del with wrong number of arguments", "DEL key1 key2", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compute.Execute(tt.command)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestInMemoryEngine(t *testing.T) {
	engine := storage.NewInMemoryEngine()

	t.Run("Set and Get", func(t *testing.T) {
		engine.Set("test_key", "test_value")
		val, ok := engine.Get("test_key")
		assert.True(t, ok)
		assert.Equal(t, "test_value", val)
	})

	t.Run("Get non-existing key", func(t *testing.T) {
		_, ok := engine.Get("non_existing_key")
		assert.False(t, ok)
	})

	t.Run("Delete", func(t *testing.T) {
		engine.Set("test_key", "test_value")
		engine.Delete("test_key")
		_, ok := engine.Get("test_key")
		assert.False(t, ok)
	})
}

func BenchmarkSet(b *testing.B) {
	log, _ := logger.NewLogger("info", "stdout")
	parser := NewParser()
	engine := storage.NewInMemoryEngine()
	compute := NewComputeLayer(parser, engine, log)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = compute.Execute("SET key value")
	}
}

func BenchmarkGet(b *testing.B) {
	log, _ := logger.NewLogger("info", "stdout")
	parser := NewParser()
	engine := storage.NewInMemoryEngine()
	compute := NewComputeLayer(parser, engine, log)

	_, _ = compute.Execute("SET key value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = compute.Execute("GET key")
	}
}

func BenchmarkDel(b *testing.B) {
	log, _ := logger.NewLogger("info", "stdout")
	parser := NewParser()
	engine := storage.NewInMemoryEngine()
	compute := NewComputeLayer(parser, engine, log)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = compute.Execute("SET key value")
		_, _ = compute.Execute("DEL key")
	}
}

```

# internal/compute/compute.go

```go
package compute

import (
	"atitov96/walgodb/internal/storage"
	"errors"
	"go.uber.org/zap"
	"sync/atomic"
	"time"
)

type Command struct {
	Type string
	Args []string
}

type Parser interface {
	Parse(expression string) (Command, error)
}

type Metrics struct {
	TotalQueries   uint64
	SuccessQueries uint64
	FailedQueries  uint64
	AverageLatency time.Duration
}

type Compute interface {
	Execute(expression string) (string, error)
	GetMetrics() Metrics
}

type computeLayer struct {
	parser  Parser
	storage storage.Storage
	log     *zap.Logger
	metrics Metrics
}

func NewComputeLayer(p Parser, s storage.Storage, l *zap.Logger) Compute {
	return &computeLayer{
		parser:  p,
		storage: s,
		log:     l,
	}
}

func (c *computeLayer) Execute(expression string) (string, error) {
	start := time.Now()
	atomic.AddUint64(&c.metrics.TotalQueries, 1)

	command, err := c.parser.Parse(expression)
	if err != nil {
		c.log.Error("failed to parse expression", zap.Error(err))
		atomic.AddUint64(&c.metrics.FailedQueries, 1)
		return "", err
	}

	c.log.Info("executing command", zap.String("command type", command.Type), zap.Strings("command args", command.Args))

	var result string
	switch command.Type {
	case "SET":
		result, err = c.handleSet(command.Args)
	case "GET":
		result, err = c.handleGet(command.Args)
	case "DEL":
		result, err = c.handleDel(command.Args)
	default:
		err = errors.New("unknown command")
	}

	if err != nil {
		atomic.AddUint64(&c.metrics.FailedQueries, 1)
	} else {
		atomic.AddUint64(&c.metrics.SuccessQueries, 1)
	}

	duration := time.Since(start)
	atomic.StoreInt64((*int64)(&c.metrics.AverageLatency), int64((c.metrics.AverageLatency*time.Duration(c.metrics.TotalQueries-1)+duration)/time.Duration(c.metrics.TotalQueries)))

	return result, err
}

func (c *computeLayer) GetMetrics() Metrics {
	return c.metrics
}

func (c *computeLayer) handleSet(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("SET requires 2 arguments")
	}
	c.storage.Set(args[0], args[1])
	return "OK", nil
}

func (c *computeLayer) handleGet(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("GET requires 1 argument")
	}
	val, ok := c.storage.Get(args[0])
	if !ok {
		return "NOT FOUND", nil
	}
	return val, nil
}

func (c *computeLayer) handleDel(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("DEL requires 1 argument")
	}
	c.storage.Delete(args[0])
	return "OK", nil
}

```

# cmd/server/main_test.go

```go
package main

import (
	"atitov96/walgodb/internal/compute"
	"atitov96/walgodb/internal/storage"
	"atitov96/walgodb/pkg/config"
	"atitov96/walgodb/pkg/logger"
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	cfg := &config.Config{
		Network: config.NetworkConfig{
			Address:        "localhost:0",
			MaxConnections: 10,
			IdleTimeout:    time.Second * 5,
		},
	}

	log, _ := logger.NewLogger("info", "stdout")
	parser := compute.NewParser()
	engine := storage.NewInMemoryEngine()
	computeLayer := compute.NewComputeLayer(parser, engine, log)

	listener, err := net.Listen("tcp", cfg.Network.Address)
	if err != nil {
		t.Fatalf("Failed to start test server: %v", err)
	}

	server := &Server{
		listener:       listener,
		computeLayer:   computeLayer,
		log:            log,
		maxConnections: cfg.Network.MaxConnections,
		idleTimeout:    cfg.Network.IdleTimeout,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go server.Serve(ctx)

	// Test connection and basic operation
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}

	_, err = fmt.Fprintf(conn, "SET test_key test_value\n")
	if err != nil {
		t.Fatalf("Failed to send SET command: %v", err)
	}

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	if string(response[:n]) != "OK\n" {
		t.Errorf("Unexpected response for SET command: %s", string(response[:n]))
	}

	// Test connection limit
	connections := make([]net.Conn, cfg.Network.MaxConnections)
	for i := 0; i < cfg.Network.MaxConnections; i++ {
		connections[i], err = net.Dial("tcp", listener.Addr().String())
		if err != nil {
			t.Fatalf("Failed to open connection %d: %v", i, err)
		}
	}

	_, err = net.Dial("tcp", listener.Addr().String())
	if err == nil {
		t.Errorf("Expected connection to be refused, but it was accepted")
	}

	// Close connections
	for _, conn := range connections {
		conn.Close()
	}

	// Test idle timeout
	conn, _ = net.Dial("tcp", listener.Addr().String())
	time.Sleep(cfg.Network.IdleTimeout + time.Second)
	_, err = fmt.Fprintf(conn, "GET test_key\n")
	//if err == nil {
	//	t.Errorf("Expected error due to idle timeout, but got none")
	//}

	cancel()
	server.Shutdown()
}

```

# cmd/server/main.go

```go
package main

import (
	"atitov96/walgodb/internal/compute"
	"atitov96/walgodb/internal/storage"
	"atitov96/walgodb/pkg/config"
	"atitov96/walgodb/pkg/logger"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	listener       net.Listener
	computeLayer   compute.Compute
	log            *zap.Logger
	maxConnections int
	idleTimeout    time.Duration
	connections    sync.WaitGroup
	semaphore      chan struct{}
}

func (s *Server) Serve(ctx context.Context) {
	s.semaphore = make(chan struct{}, s.maxConnections)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				s.log.Error("Error accepting connections", zap.Error(err))
			}
		}

		select {
		case s.semaphore <- struct{}{}:
			s.connections.Add(1)
			go s.handleConnection(ctx, conn)
		default:
			s.log.Error("Max connections reached, closed new connection")
			conn.Close()
		}
	}
}

func (s *Server) Shutdown() {
	s.listener.Close()
	s.connections.Wait()
}

func (s *Server) handleConnection(ctx context.Context, conn net.Conn) {
	defer func() {
		conn.Close()
		<-s.semaphore
		s.connections.Done()
	}()

	for {
		err := conn.SetDeadline(time.Now().Add(s.idleTimeout))
		if err != nil {
			s.log.Error("Error setting connection deadline", zap.Error(err))
			return
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				s.log.Info("Connection idle timeout")
			} else {
				s.log.Error("Error reading from connection", zap.Error(err))
			}
			return
		}

		query := string(buffer[:n])
		result, err := s.computeLayer.Execute(query)
		if err != nil {
			s.log.Error("Error executing query", zap.Error(err))
			_, writeErr := conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			if writeErr != nil {
				s.log.Error("Error writing error response", zap.Error(writeErr))
				return
			}
		} else {
			_, writeErr := conn.Write([]byte(fmt.Sprintf("%s\n", result)))
			if writeErr != nil {
				s.log.Error("Error writing response", zap.Error(writeErr))
				return
			}
		}

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	log, err := logger.NewLogger(cfg.Logging.Level, cfg.Logging.Output)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	parser := compute.NewParser()
	engine := storage.NewInMemoryEngine()
	computeLayer := compute.NewComputeLayer(parser, engine, log)

	listener, err := net.Listen("tcp", cfg.Network.Address)
	if err != nil {
		log.Fatal("Failed to start TCP server", zap.Error(err))
	}

	server := &Server{
		listener:       listener,
		computeLayer:   computeLayer,
		log:            log,
		maxConnections: cfg.Network.MaxConnections,
		idleTimeout:    cfg.Network.IdleTimeout,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go server.Serve(ctx)

	log.Info("TCP server started", zap.String("address", cfg.Network.Address))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	log.Info("Shutting down gracefully...")
	cancel()
	server.Shutdown()
}

```

# cmd/client/main.go

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	address := flag.String("address", "localhost:3223", "Address of the database server")
	timeout := flag.Duration("timeout", 5*time.Second, "Timeout for connection and operations")
	flag.Parse()

	conn, err := net.DialTimeout("tcp", *address, *timeout)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to WalGoDb server. Type 'exit' to quit or 'help' for help.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("WalGoDb> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}
		if strings.ToLower(input) == "help" {
			printHelp()
			continue
		}

		_, err := fmt.Fprintf(conn, "%s\n", input)
		if err != nil {
			fmt.Printf("Error sending to server: %v\n", err)
			continue
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from server: %v\n", err)
			continue
		}

		fmt.Print(response)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}

func printHelp() {
	fmt.Println(`
Available commands:
SET <key> <value> : Set a key-value pair
GET <key>         : Get the value for a key
DEL <key>         : Delete a key-value pair
help              : Show this help message
exit              : Exit the program

Note: Keys and values must be alphanumeric (including underscores and hyphens)
	`)
}

```

