include $(GOROOT)/src/Make.inc
PKGDIR=$(GOROOT)/pkg/$(GOOS)_$(GOARCH)



DIRS=\
	gmf



all: deps

clean: myclean


deps : 
	for dir in $(DIRS); do (cd $$dir; $(MAKE) install ); done 
#	cd pkg; $(MAKE) $(MFLAGS); $(MAKE) install \
#	cd ffmpeg; $(MAKE) $(MFLAGS); $(MAKE) install


myclean :
	for d in $(DIRS); do (cd $$d; $(MAKE) clean ); done 
#	echo cleaning up in .
#	-$(RM) -f $(EXE) $(OBJS) $(OBJLIBS)
#	for d in pkg; do (cd $$d; $(MAKE) clean ); done 
#	for d in ffmpeg; do (cd $$d; $(MAKE) clean ); done
test :
	for d in $(DIRS); do (cd $$d; $(MAKE) test ); done 

include $(GOROOT)/src/Make.cmd


