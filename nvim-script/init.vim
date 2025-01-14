" General Settings
 set number
 set relativenumber
 set autoindent
 set paste
 set tabstop=4
 set shiftwidth=4
 set smarttab
 set softtabstop=4
 set encoding=UTF-8
 set completeopt=menuone,noinsert,noselect

" Telescope mappings
nnoremap <leader>ff <cmd>Telescope find_files<cr>
nnoremap <leader>fg <cmd>Telescope live_grep<cr>
nnoremap <leader>fb <cmd>Telescope buffers<cr>
nnoremap <leader>fh <cmd>Telescope help_tags<cr>

 " Plugin Manager
 call plug#begin()

 " Plugins
 Plug 'tpope/vim-surround' " Surrounding ysw)
 Plug 'preservim/nerdtree' " NerdTree
 Plug 'tpope/vim-commentary' " For Commenting gcc & gc
 Plug 'vim-airline/vim-airline' " Status bar
 Plug 'ap/vim-css-color' " CSS Color Preview
 Plug 'rafi/awesome-vim-colorschemes' " Retro Scheme
 Plug 'neoclide/coc.nvim', {'branch': 'release'} " Auto Completion
 Plug 'ryanoasis/vim-devicons' " Developer Icons
 Plug 'tc50cal/vim-terminal' " Vim Terminal
 Plug 'preservim/tagbar' " Tagbar for code navigation
 Plug 'terryma/vim-multiple-cursors' " CTRL + N for multiple cursors
 Plug 'nvim-tree/nvim-web-devicons'
 Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
Plug 'wakatime/vim-wakatime'

 call plug#end()

 " Key Mappings
 nnoremap <C-f> :NERDTreeFocus<CR>
 nnoremap <C-n> :NERDTree<CR>
 nnoremap <C-t> :NERDTreeToggle<CR>

 " Plugin-Specific Configurations
 let g:airline_powerline_fonts = 1

 " vim-go Settings (if used)
 let g:go_code_completion_enabled = 0 " Disable if using coc.nvim
 let g:go_def_mode='gopls'
 let g:go_info_mode='gopls'

 " coc.nvim Settings
 autocmd BufWritePre *.go :call CocAction('runCommand', 'editor.action.organizeImport')

 " Colorscheme
 colorscheme materialbox
