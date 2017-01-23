# Package configuration
PROJECT = framework 
COMMANDS =
CODECOV_TOKEN = 59e8b601-724f-46f1-8380-7b917d5cc704 

# Including devops Makefile
MAKEFILE = Makefile.main
DEVOPS_REPOSITORY = https://github.com/src-d/devops.git
DEVOPS_FOLDER = .devops
CI_FOLDER = .ci

$(MAKEFILE):
	@git clone --quiet $(DEVOPS_REPOSITORY) $(DEVOPS_FOLDER); \
	cp -r $(DEVOPS_FOLDER)/ci .ci; \
	rm -rf $(DEVOPS_FOLDER); \
	cp $(CI_FOLDER)/$(MAKEFILE) .;

-include $(MAKEFILE)
