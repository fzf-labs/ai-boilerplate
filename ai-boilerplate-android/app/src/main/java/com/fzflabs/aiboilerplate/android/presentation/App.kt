package com.fzflabs.aiboilerplate.android.presentation

import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.hilt.lifecycle.viewmodel.compose.hiltViewModel
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.fzflabs.aiboilerplate.android.presentation.screen.HomeScreen
import com.fzflabs.aiboilerplate.android.presentation.viewmodel.HomeViewModel

@Composable
fun App() {
    val navController = rememberNavController()
    val viewModel = hiltViewModel<HomeViewModel>()

    NavHost(navController = navController, startDestination = "home") {
        composable("home") {
            HomeScreen(viewModel.productDetailsState.collectAsState().value)
        }
    }
}
