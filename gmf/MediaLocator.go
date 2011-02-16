package gmf

import "strings"

type MediaLocator struct{
  Filename string;
  Format string
}

func (loc * MediaLocator) GetProtocol() string{
  lines:=strings.Split(loc.Filename,":",-1)
  if(len(lines)!=2){
    return "file"
  }
  return lines[0]
}

func (loc * MediaLocator) GetReminder()string{
  lines:=strings.Split(loc.Filename,":",-1)
  if(len(lines)!=2){
    return loc.Filename
  }
  return lines[1][2:]
}
