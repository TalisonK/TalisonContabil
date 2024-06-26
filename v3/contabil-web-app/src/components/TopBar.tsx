import React, { useState, useEffect } from 'react'
import { Logo, SidebarContainer } from '../pages/dashboard/styles'
import {
    AppBar,
    Avatar,
    Box,
    Button,
    Container,
    IconButton,
    Menu,
    MenuItem,
    Toolbar,
    Tooltip,
    Typography,
} from '@mui/material'
import { ThemeSwitch } from './styles'
import { DisplayFlex } from '../styles'
import User from '../interfaces/User'

interface TopBarProps {
    theme: string
    setTheme: (theme: boolean) => void
}

const TopBar = (props: TopBarProps) => {
    const pages = [
        {name: 'Insert', role: 'ROLE_USER'}, 
        {name: 'Activities', role: 'ROLE_USER'},
        {name: 'Categories', role: 'ROLE_USER'},
        {name: 'Users', role: 'ROLE_ADMIN'}
    ]
    const settings = ['Profile', 'Logout']

    const [user, setUser] = useState<User>({id:"",name:"A",role:'ROLE_USER',createdAt:""})

    useEffect(() => {
        const userStorage = localStorage.getItem('user')
        if (userStorage) {
            setUser(JSON.parse(userStorage))
        }
    }, [])

    const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(
        null
    )

    const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorElUser(event.currentTarget)
    }

    const handleCloseUserMenu = () => {
        setAnchorElUser(null)
    }

    const handleLogout = (setting: string) => {
        if (setting === 'Logout') {
            localStorage.removeItem('user')
            window.location.href = '/'
        }
    }

    return (
        <AppBar
            position="static"
            id="sidebar"
            style={{
                backgroundColor: props.theme === 'dark' ? 'rgb(16 15 47)' : '',
            }}
        >
            <Container maxWidth="xl">
                <Toolbar disableGutters>
                    <Typography
                        variant="h6"
                        noWrap
                        component="a"
                        href="/"
                        sx={{ mr: 2, display: { xs: 'none', md: 'flex' } }}
                    >
                        <Logo src="logo-invertido-lateral.png" alt="logo" />
                    </Typography>

                    <Typography
                        variant="h5"
                        noWrap
                        component="a"
                        href="#app-bar-with-responsive-menu"
                        sx={{
                            mr: 2,
                            display: { xs: 'flex', md: 'none' },
                            flexGrow: 1,
                        }}
                    >
                        <Logo src="logo-invertido-lateral.png" alt="logo" />
                    </Typography>
                    <Box
                        sx={{
                            flexGrow: 1,
                            display: { xs: 'none', md: 'flex' },
                        }}
                    >
                        {pages.map((page) => {

                            if (page.role === 'ROLE_ADMIN') {
                                return user.role == 'ROLE_ADMIN'?
                                (
                                <Button
                                    key={page.name}
                                    href={`/${page.name}`}
                                    sx={{
                                        my: 2,
                                        color: 'white',
                                        display: 'block',
                                    }}
                                >
                                    {page.name}
                                </Button>
                                ):
                                <></>
                            } else {
                                return (<Button
                                    key={page.name}
                                    href={`/${page.name}`}
                                    sx={{
                                        my: 2,
                                        color: 'white',
                                        display: 'block',
                                    }}
                                >
                                    {page.name}
                                </Button>)
                            }
                        })}
                    </Box>

                    <Box sx={{ flexGrow: 0 }}>
                        <Tooltip title="Open settings">
                            <IconButton
                                onClick={handleOpenUserMenu}
                                sx={{ p: 0 }}
                            >
                                <Avatar
                                    alt={`${user.name.toUpperCase()}`}
                                    src="/static/images/avatar/2.jpg"
                                />
                            </IconButton>
                        </Tooltip>
                        <Menu
                            sx={{ mt: '45px' }}
                            id="menu-appbar"
                            anchorEl={anchorElUser}
                            anchorOrigin={{
                                vertical: 'top',
                                horizontal: 'right',
                            }}
                            keepMounted
                            transformOrigin={{
                                vertical: 'top',
                                horizontal: 'right',
                            }}
                            open={Boolean(anchorElUser)}
                            onClose={handleCloseUserMenu}
                        >
                            <DisplayFlex justifyContent="center">
                                <ThemeSwitch
                                    checked={props.theme === 'dark'}
                                    onChange={(e) => {
                                        props.setTheme(e.target.checked)
                                    }}
                                />
                            </DisplayFlex>
                            {settings.map((setting) => (
                                <MenuItem
                                    key={setting}
                                    onClick={() => handleLogout(setting)}
                                >
                                    <Typography textAlign="center">
                                        {setting}
                                    </Typography>
                                </MenuItem>
                            ))}
                        </Menu>
                    </Box>
                </Toolbar>
            </Container>
        </AppBar>
    )
}

export default TopBar
