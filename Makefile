# Define the name of the executable file
TARGET_NAME := Tally Connector
# Define the source file
SOURCE_FILE := main.go

# --- Platform-specific variables ---fdg

# Define the linker flags. 
# The -H=windowsgui flag tells the linker to create a GUI application 
# for Windows, suppressing the console window.
# We use the shell 'cmd' to check the GOOS environment variable.
LDFLAGS := 

ifeq ($(OS),Windows_NT)
    # Windows: Use the GUI flag and the .exe extension
    LDFLAGS := -ldflags "-H=windowsgui"
    OUTPUT_FILE := $(TARGET_NAME).exe
else
    # macOS/Linux: No special flags, no file extension
    OUTPUT_FILE := $(TARGET_NAME)
endif

# --- Main Build Target ---

.PHONY: build

build:
	@echo "Building $(OUTPUT_FILE) for $(shell go env GOOS)..."
	go build $(LDFLAGS) -o $(OUTPUT_FILE) $(SOURCE_FILE)

# --- Clean Target ---

.PHONY: clean

clean:
	@echo "Cleaning up..."
	@if [ -f "$(TARGET_NAME)" ]; then rm -f "$(TARGET_NAME)"; fi
	@if [ -f "$(TARGET_NAME).exe" ]; then rm -f "$(TARGET_NAME).exe"; fi
	@echo "Cleanup complete."