package main

import "os"

// flag value for existing file
type fileValue struct{
    *os.File
}

func (fsf *fileValue) Set(v string) (err error) {
    fsf.File,err=os.Open(v)
    return
}

func (fsf *fileValue) String() string {
    if fsf==nil || fsf.File==nil {return "<nil>"}
    return fsf.File.Name()
}


// flag value for file, creates if needed.
type createFileValue struct{
    fileValue
}

func (fsf *createFileValue) Set(v string) (err error) {
	fsf.File,err=os.Create(v)
	return
}

// flag value for an existing directory.
type dirValue struct{
    fileValue
}

func (fsf *dirValue) Set(v string) error {
	f,err:=os.Open(v)
    if err!=nil{ return err}
    fi,err:=f.Stat()
    if err!=nil{ return err}
	if !fi.IsDir(){return os.ErrNotExist}
    fsf.File=f
    return err
}

// flag value for an existing directory, creates if needed.
type newDirValue struct{
    fileValue
}

func (fsf *newDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
	if os.IsNotExist(err){
		err=os.Mkdir(v,0777)
 		if err!=nil{ return}
		f,err=os.Open(v)
 		if err!=nil{ return}
 		fsf.File=f
		return
	}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
    fsf.File=f
    return
}

// flag value for a directory, creates if needed, any pre-existing hierarchy inside it is erased.
type newOverwriteDirValue struct{
    fileValue
}

func (fsf *newOverwriteDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
	if os.IsNotExist(err){
		err=os.Mkdir(v,0777)
 		if err!=nil{ return}
		f,err=os.Open(v)
 		if err!=nil{ return}
 		fsf.File=f
		return
	}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=RemoveContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, pre-existing, emptied.
type overwriteDirValue struct{
    fileValue
}

func (fsf *overwriteDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
    if err!=nil{ return}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=RemoveContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, creates if needed, any pre-existing files inside it are erased, (but not directories).
type newOverwriteFilesDirValue struct{
    fileValue
}

func (fsf *newOverwriteFilesDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
	if os.IsNotExist(err){
		err=os.Mkdir(v,0777)
 		if err!=nil{ return}
		f,err=os.Open(v)
 		if err!=nil{ return}
 		fsf.File=f
		return
	}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=RemoveFileContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, pre-existing, any pre-existing files inside it are erased, (but not directories).
type overwriteFilesDirValue struct{
    fileValue
}

func (fsf *overwriteFilesDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
    if err!=nil{ return}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=RemoveFileContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}


// flag value for a directory, creates if needed, any pre-existing directories inside it are erased, (but not files).
type newOverwriteSubdirsDirValue struct{
    fileValue
}

func (fsf *newOverwriteSubdirsDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
	if os.IsNotExist(err){
		err=os.Mkdir(v,0777)
 		if err!=nil{ return}
		f,err=os.Open(v)
 		if err!=nil{ return}
 		fsf.File=f
		return
	}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=RemoveDirContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, pre-existing, any pre-existing directories inside it are erased, (but not files).
type overwriteSubdirsDirValue struct{
    fileValue
}

func (fsf *overwriteSubdirsDirValue) Set(v string) (err error) {
	f,err:=os.Open(v)
    if err!=nil{ return}
    fi,err:=f.Stat()
    if err!=nil{ return}
	if !fi.IsDir(){return os.ErrNotExist}
	err=RemoveDirContents(f)
    if err!=nil{ return}
    fsf.File=f
    return
}

// flag value for a directory, not pre-existing.
type makeDirValue struct{
    fileValue
}

func (fsf *makeDirValue) Set(v string) (err error) {
	err=os.Mkdir(v,0777)
    if err!=nil{ return}
    fsf.File,err=os.Open(v)
    return
}

// flag value for a directory, not pre-existing, possibly multiple levels down.
type makeDirAllValue struct{
    fileValue
}

func (fsf *makeDirAllValue) Set(v string) (err error) {
	err=os.MkdirAll(v,0777)
    if err!=nil{ return}
    fsf.File,err=os.Open(v)
   return
}

// flag value for a directory, possibly down multiple levels. if pre-existing erased.
type makeDirOverwriteAllValue struct{
    fileValue
}

func (fsf *makeDirOverwriteAllValue) Set(v string) (err error) {
	err=os.RemoveAll(v)
    if err!=nil{ return}
	err=os.MkdirAll(v,0777)
    fsf.File,err=os.Open(v)
   return
}

// flag value for a new directory at this level. if pre-existing erased.
type makeDirOverwriteValue struct{
    fileValue
}

func (fsf *makeDirOverwriteValue) Set(v string) (err error) {
	err=os.RemoveAll(v)
    if err!=nil{ return}
	err=os.Mkdir(v,0777)
    fsf.File,err=os.Open(v)
   return
}

func RemoveContents(d *os.File) error {
	finfos, err := d.Readdir(-1)
    if err != nil {
        return err
    }
	defer  changeWorkingDirReset(d)()
    for _, finfo := range finfos {
    	if finfo.IsDir(){
			err=os.RemoveAll(finfo.Name())	
    	}else{	
			err=os.Remove(finfo.Name())	
		}
	    if err != nil {
	        return err
	    }
    }
    return nil
}

func RemoveFileContents(d *os.File) error {
	finfos, err := d.Readdir(-1)
    if err != nil {
        return err
    }
	defer  changeWorkingDirReset(d)()
    for _, finfo := range finfos {
    	if !finfo.IsDir(){
			err=os.Remove(finfo.Name())	
		}
	    if err != nil {
	        return err
	    }
    }
    return nil
}

func RemoveDirContents(d *os.File) error {
	finfos, err := d.Readdir(-1)
    if err != nil {
        return err
    }
	defer  changeWorkingDirReset(d)()
    for _, finfo := range finfos {
    	if finfo.IsDir(){
			err=os.RemoveAll(finfo.Name())	
		}
	    if err != nil {
	        return err
	    }
    }
    return nil
}


func changeWorkingDirReset(dir *os.File) (fn func()) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = dir.Chdir()
	if err == nil {
		return func() { os.Chdir(currentDir) }
	}
	return
}
