import { Alert, Box, Divider, FormControl, Grid, Select, Snackbar, TextField, Typography } from "@mui/material";
import React, { useEffect, useState } from "react";
import Button from "@mui/material/Button";
// import {Link as RouterLink} from "react-router-dom"
// import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';



  function DrugCreate(){
    return(
        <Box sx={{
            backgroundImage:
              "url(https://i.pinimg.com/564x/c0/4a/5f/c04a5f27b605cc21aa55420a7f83318d.jpg)",
            backgroundRepeat: "no-repeat",
            backgroundSize: "cover",
            width: "100%",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}>
      
            <Box
              sx={{
                fontFamily: "PK Krung Thep Medium",
                fontSize: "20px",
                width: "90%",
                mt: "70px",
                bgcolor: 'rgba(255, 255, 255, 0.5)',
                borderRadius: "30px",
                boxShadow: 20,
      
              }}
            >
      
              <Snackbar open={false} autoHideDuration={6000} onClose={()=>{}}>
                <Alert onClose={() => {}} severity="success">
                  บันทึกข้อมูลสำเร็จ
                </Alert>
              </Snackbar>
              <Snackbar open={false} autoHideDuration={6000} onClose={() => {}}>
                <Alert onClose={() => {}} severity="error">
                  บันทึกข้อมูลไม่สำเร็จ: {}
                </Alert>
              </Snackbar>
      
              <Box sx={{ padding: 2, color: "text.secondary" }}>
                <Box display="flex">
                  <Box flexGrow={1}>
                    <Typography
                      component="h6"
                      variant="h5"
                      color="primary"
                      gutterBottom
                      sx={{
                        display: 'flex',
                        justifyContent: 'center',
                        
                      }}
                    >
                      <h1>บันทึกการชำระเงิน</h1>
      
                    </Typography>
                  </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ flexGrow: 1 }}>
      
      
                  <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <p>ชื่อ - นามสกุล</p>
                      <Select
                        native
                        disabled
                        value={[]}
                        onChange={() => {}}
                        inputProps={{
                          name: "UserID",
                        }}
                      >
                        <option value={[]} key={0} >
                          {/* {users?.Name} */}
                        </option>
                      </Select>
                    </FormControl>
                  </Grid>
      
      
      
                  <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <p>เลขห้อง</p>
                      <TextField
                        variant="outlined"
                        disabled
                        id="BillID"
                        inputProps={{
                          name: "BillID",
                        }}
      
                        value={[]}
                      />
      
                    </FormControl>
                  </Grid>
      
      
                  <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <p>ค่าห้องพัก</p>
                      <TextField
                        variant="outlined"
                        disabled
                        id="BillID"
                        inputProps={{
                          name: "BillID",
                        }}
      
                        value={[]}
                      />
                    </FormControl>
                  </Grid>
      
      
                  <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <p>ค่าน้ำ</p>
                      <TextField
                        variant="outlined"
                        disabled
                        id="BillID"
                        inputProps={{
                          name: "BillID",
                        }}
      
                        value={[]}
                      />
                    </FormControl>
                  </Grid>
      
      
                  <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <p>ค่าไฟ</p>
                      <TextField
                        variant="outlined"
                        disabled
                        id="BillID"
                        inputProps={{
                          name: "BillID",
                        }}
      
                        value={[]}
                      />
                    </FormControl>
                  </Grid>
      
      
                  <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <p>ค่าเฟอร์นิเจอร์</p>
                      <TextField
                        variant="outlined"
                        disabled
                        id="BillID"
                        inputProps={{
                          name: "BillID",
                        }}
      
                        value={[]}
                      />
                    </FormControl>
                  </Grid>
      
                  <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <p>ยอดรวมชำระ</p>
                      <TextField
                        variant="outlined"
                        disabled
                        id="BillID"
                        inputProps={{
                          name: "BillID",
                        }}
      
                        value={[]}
                      />
                    </FormControl>
                  </Grid>
      
      
      
                  <Grid item xs={6} sx={{ mt: 'auto', }}>
                    <FormControl fullWidth variant="outlined">
                      <p>ช่องทางการชำระ</p>
                      <Select
                        native
                        value={[]}
                        onChange={() => {}}
                        inputProps={{
                          name: "MethodID",
                        }}
                      >
                        <option aria-label="None" value="">
                          เลือก...
                        </option>
                        {/* {methods.map((item: MethodInterface) => (
                          <option value={item.ID} key={item.ID}>
                            {item.Name}
                          </option>
                        ))} */}
                      </Select>
                    </FormControl>
                  </Grid>
      
      
                  <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <p>ธนาคาร</p>
                      <Select
                        native
                        value={[]}
                        onChange={() => {}}
                        inputProps={{
                          name: "BankingID",
                        }}
                      >
                        <option aria-label="None" value="">
                          เลือก...
                        </option>
                        {/* {bankings.map((item: BankingInterface) => (
                          <option value={item.ID} key={item.ID}>
                            {item.Name}
                          </option>
                        ))} */}
                      </Select>
                    </FormControl>
                  </Grid>
      
      
                  {/* <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <LocalizationProvider dateAdapter={AdapterDateFns}>
                        <p>ช่องทางการชำระ</p>
                        <DateTimePicker
                          value={selectedDate}
                          onChange={(newValue) => setSelectedDate(newValue)}
                          minDate={(new Date)}   //บันทึกค่าปัจจุบัน
                          renderInput={(params) =>
                            <TextField {...params} />}
                        />
                      </LocalizationProvider>
                    </FormControl>
                  </Grid> */}
      
      
                  <Grid item xs={6}>
                    <FormControl fullWidth variant="outlined">
                      <Button variant="contained" component="label">
                        Upload
                        <input hidden accept="image/*" multiple type="file" onChange={() => {}} />
                      </Button>
                    </FormControl>
                  </Grid>
      
                  <Grid item xs={12}>
                    <Button
                    //   component={RouterLink}
                    //   to="/payments"
                    //   variant="contained"
                    >
                      กลับ
                    </Button>
                    <Button
                    //   style={{ float: "right" }}
                    //   variant="contained"
                    //   onClick={[]}
                    //   color="primary"
                    >
                      บันทึกข้อมูล
                    </Button>
                  </Grid>
                </Grid>
              </Box>
            </Box>
          </Box>
    )
  }
  export default DrugCreate