EMACS ?= emacs
BATCH = $(EMACS) -Q --batch -L .
TESTS = $(if $(DAY),day$(DAY)-test.el,$(wildcard day*-test.el))

.PHONY: test run

test:
	$(BATCH) $(addprefix -l ,$(TESTS)) -f ert-run-tests-batch-and-exit

run:
	@test -n "$(DAY)" || (echo "usage: make run DAY=01" && exit 1)
	@$(BATCH) -l day$(DAY).el --eval "(aoc-run $(DAY) #'day$(DAY)-part1 #'day$(DAY)-part2)"
