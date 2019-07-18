"Auto generate ctags for project with a base dir containing '.git' directory
"and will place it inside there (so we aren't committing a tags file)

"Projects without a tags file generated inside will have one generated
"and those _with_ one will have their tags file appended to

"Due to lack of async in vim, we just run ctags in a background job

let globaldir=getcwd(-1)
let current_tags_file=(getcwd(-1) . ".git/tags")

function TagsExists()
  if isdirectory(getcwd(-1) + ".git")
    if !filereadable(getcwd(-1) + "/.git/tags")
      return 0
    endif
  endif
  return 1
endfunction

function CreateTagsFile()
  if !TagsExists()
    system("touch",getcwd(-1) + "/.git/tags")

    let tagscmd=("ctags -f " . "'" . getcwd(-1) ."'" . "/.git/tags" . " -R * &")
    return system(tagscm)
    "WriteNotice("Created tags file: " + getcwd(-1) + "/.git/tags")
  endif
endfunction

function UpdateTags()
  let tagscmd=("ctags -f " . "'" . getcwd(-1) ."'" . "/.git/tags" . " --append * &")
  if !TagsExists()
    CreateTagsFile()
  else
    return system(tagscmd)
  end
endfunction

set tags=current_tags_file

autocmd BufWrite,BufEnter * call UpdateTags()
